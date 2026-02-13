package gorms

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db      *gorm.DB // Default DB for backward compatibility
	clients = make(map[string]*gorm.DB)
)

func InitFormViper() (*gorm.DB, error) {
	viper.SetDefault("DB_DSN", "")
	viper.SetDefault("DB_MAX_OPEN_CONNS", 100)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 20)

	// Initialize multiple clients if configured
	if viper.IsSet("db.client") {
		configs := viper.GetStringMap("db.client")
		for name := range configs {
			dialect := viper.GetString(fmt.Sprintf("db.client.%s.dialect", name))
			dsn := viper.GetString(fmt.Sprintf("db.client.%s.dsn", name))

			client, err := NewDatabase(dialect, dsn, &gorm.Config{})
			if err != nil {
				return nil, err
			}

			sqlDB, err := client.DB()
			if err != nil {
				return nil, err
			}

			sqlDB.SetMaxOpenConns(viper.GetInt("DB_MAX_OPEN_CONNS"))
			sqlDB.SetMaxIdleConns(viper.GetInt("DB_MAX_IDLE_CONNS"))
			sqlDB.SetConnMaxIdleTime(time.Second * 5)
			sqlDB.SetConnMaxLifetime(time.Hour)

			clients[name] = client
		}

		if client, ok := clients["default"]; ok {
			db = client
		}
	}

	// Legacy support: if db is not set (either no "default" in client map, or no client map at all)
	// try reading from root db.dialect and db.dsn
	if db == nil {
		dialect := viper.GetString("db.dialect")
		dsn := viper.GetString("db.dsn")

		if dialect != "" && dsn != "" {
			var err error
			db, err = NewDatabase(dialect, dsn, &gorm.Config{})
			if err != nil {
				return nil, err
			}

			sqlDB, err := db.DB()
			if err != nil {
				return nil, err
			}

			sqlDB.SetMaxOpenConns(viper.GetInt("DB_MAX_OPEN_CONNS"))
			sqlDB.SetMaxIdleConns(viper.GetInt("DB_MAX_IDLE_CONNS"))
			sqlDB.SetConnMaxIdleTime(time.Second * 5)
			sqlDB.SetConnMaxLifetime(time.Hour)

			clients["default"] = db
		}
	}

	return db, nil
}

func InitGenFromViper(setDefault func(db *gorm.DB, opts ...gen.DOOption)) error {
	var err error
	if db, err = InitFormViper(); err != nil {
		return err
	}

	// Apply debug/logger settings to all clients
	for _, client := range clients {
		if viper.GetBool("debug") {
			client = client.Debug()
		}
		if viper.GetBool("local") {
			client.Config.Logger = logger.Default.LogMode(logger.Silent)
		}
	}

	// Also apply to the default db reference if it exists (it should point to one of the clients or be independent)
	// Since 'db' is a pointer, modifying 'clients["default"]' should reflect if they point to same struct?
	// Wait, db = db.Debug() returns a NEW session. It doesn't modify the global *gorm.DB in place in a way that affects other references unless we update the variable.
	// However, InitGenFromViper is usually called once at startup.

	if db != nil {
		if viper.GetBool("debug") {
			db = db.Debug()
			// We should probably update the map too if we want consistency, but 'db' is the main return value here.
			if _, ok := clients["default"]; ok {
				clients["default"] = db
			}
		}
		if viper.GetBool("local") {
			db.Config.Logger = logger.Default.LogMode(logger.Silent)
		}
		setDefault(db)
	}

	return nil
}

func Client() *gorm.DB {
	return db
}

func GetClient(name string) *gorm.DB {
	if client, ok := clients[name]; ok {
		return client
	}
	return nil
}
