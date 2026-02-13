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
			g.ApplyBasic(
				model.CRMAdmin{},
				model.CRMPermission{},
				model.CRMRole{},
				model.CRMRolePermission{},
				model.CRMArticle{},
				model.CRMCategory{},
				model.CRMArticleContent{},
				model.SPictureBook{},
				model.SPictureBookItem{},
				model.SPictureBookCategory{},
			)
			g.Execute()
			return nil
		},
	}
}
