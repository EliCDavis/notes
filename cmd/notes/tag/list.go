package tag

import (
	"io"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "lists all tags",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			var out io.Writer = ctx.App.Writer
			return project.ListTags(out)
		},
	}
}
