package commands

import (
	"crm/commands/generate"
	"crm/commands/gorm"
	"crm/commands/migrate"
	"crm/commands/worker"

	"github.com/urfave/cli/v2"
)

func All() []*cli.Command {
	commands := []*cli.Command{
		migrate.Command(),
		generate.Command(),
		gorm.Command(),
		worker.Command(),
	}
	return commands
}
