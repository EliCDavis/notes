package log

import (
	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "log",
		Usage: "Create and edit logs",
		Flags: []cli.Flag{
			flags.Project,
		},
		Subcommands: []*cli.Command{
			newCommand(),
			openCommand(),
		},
	}
}
