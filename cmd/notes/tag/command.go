package tag

import (
	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "tag",
		Usage: "Manage tags",
		Flags: []cli.Flag{
			flags.Project,
		},
		Subcommands: []*cli.Command{
			newCommand(),
			updateCommand(),
			listCommand(),
		},
	}
}
