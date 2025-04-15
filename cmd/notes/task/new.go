package task

import (
	"strings"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func newCommand() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "Creates a new task",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "name",
			},
		},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			taskName := strings.TrimSpace(ctx.String("name"))
			if err = project.NewTask(taskName); err != nil {
				return err
			}

			return project.Save()
		},
	}
}
