package task

import (
	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "task",
		Usage: "Create and edit tasks",
		Flags: []cli.Flag{
			flags.Project,
		},
		Subcommands: []*cli.Command{
			newCommand(),
			listCommand(),
			updateCommand(),
			todoCommand(),
			ganttCommand(),
		},
	}
}
