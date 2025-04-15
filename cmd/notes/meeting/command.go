package meeting

import (
	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "meeting",
		Usage: "Create and edit meetings",
		Flags: []cli.Flag{
			flags.Project,
		},
		Subcommands: []*cli.Command{
			newCommand(),
		},
	}
}
