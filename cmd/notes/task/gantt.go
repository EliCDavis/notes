package task

import (
	"io"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func ganttCommand() *cli.Command {
	return &cli.Command{
		Name:  "gantt",
		Usage: "create a gantt chart in mermaid",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			var out io.Writer = ctx.App.Writer
			project.TaskGantt(out)
			return nil
		},
	}
}
