package task

import (
	"fmt"
	"io"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "lists all tasks",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			var out io.Writer = ctx.App.Writer
			for i, task := range project.Tasks {
				status := "Created"
				statusTime := task.Created
				if len(task.History) > 0 {
					item := task.History[len(task.History)-1]
					status = string(item.Status)
					statusTime = item.Time
				}

				fmt.Fprintf(out, "[%d] %-10s %s - %s\n", i, status, statusTime.Format("2006-01-02 15:04"), task.Name)
			}

			return nil
		},
	}
}
