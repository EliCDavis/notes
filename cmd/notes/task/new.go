package task

import (
	"errors"
	"strings"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func newCommand() *cli.Command {
	return &cli.Command{
		Name:      "new",
		Usage:     "Creates a new task",
		Args:      true,
		ArgsUsage: "[Task Name]",
		Flags:     []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			args := ctx.Args()
			if args.Len() > 1 {
				return errors.New("expected 0 or 1 argument for task name")
			}

			name := strings.TrimSpace(args.First())
			if err = project.NewTask(name); err != nil {
				return err
			}

			return project.Save()
		},
	}
}
