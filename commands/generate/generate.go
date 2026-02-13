package generate

import (
	"crm/gopkg/gorms"

	"github.com/urfave/cli/v2"
	"gorm.io/gen"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "generate",
		Action: func(ctx *cli.Context) error {
			g := gen.NewGenerator(gen.Config{
				OutPath: "internal/g",
				Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
			})
			g.UseDB(gorms.Client())
			g.ApplyBasic()
			g.Execute()
			return nil
		},
	}
}
