package topic

import (
	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "topic",
		Usage: "Create and edit topics",
		Flags: []cli.Flag{
			flags.Project,
		},
		Subcommands: []*cli.Command{
			newCommand(),
		},
	}
}
