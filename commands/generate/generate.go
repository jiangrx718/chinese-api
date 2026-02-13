package generate

import (
	"crm/gopkg/gorms"
	"crm/internal/model"

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
			// Generate basic type-safe DAO API for struct following conventions
			g.ApplyBasic(
				model.SBookName{},
				model.SChinesePicture{},
				model.SChinesePictureInfo{},
			)
			g.Execute()
			return nil
		},
	}
}
