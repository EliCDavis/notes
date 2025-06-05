package task

import (
	"io"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func currentCommand() *cli.Command {
	return &cli.Command{
		Name:  "current",
		Usage: "lists tasks are in progress",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			var out io.Writer = ctx.App.Writer
			return project.ListCurrentTasks(out)
		},
	}
}
