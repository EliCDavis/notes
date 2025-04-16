package task

import (
	"io"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func todoCommand() *cli.Command {
	return &cli.Command{
		Name:  "todo",
		Usage: "lists tasks that still need to be done",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			var out io.Writer = ctx.App.Writer
			return project.ListTodoTasks(out)
		},
	}
}
