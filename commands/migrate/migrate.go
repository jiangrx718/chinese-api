package migrate

import (
	"crm/gopkg/gorms"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func doSeed(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {

		return nil
	})
}

func Command() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "数据库迁移",
		Subcommands: []*cli.Command{
			{
				Name:        "up",
				Usage:       "自动迁移数据库",
				Description: "自动迁移数据库",
				Action: func(ctx *cli.Context) error {
					tx := gorms.Client()
					tx.DisableForeignKeyConstraintWhenMigrating = true
					tables := []any{}
					if err := tx.AutoMigrate(tables...); err != nil {
						return err
					}
					return doSeed(tx)
				},
			},
			{
				Name:        "seed",
				Usage:       "初始化基础数据",
				Description: "插入管理员、角色与权限等初始数据",
				Action: func(ctx *cli.Context) error {
					return doSeed(gorms.Client())
				},
			},
		},
	}
}
