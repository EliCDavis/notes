package images

import (
	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "image",
		Usage: "Manage images",
		Flags: []cli.Flag{
			flags.Project,
		},
		Subcommands: []*cli.Command{
			addCommand(),
		},
	}
}
