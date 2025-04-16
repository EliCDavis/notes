package images

import (
	"errors"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func addCommand() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     "Adds an image to the project",
		Args:      true,
		ArgsUsage: "[path to images]",
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			args := ctx.Args()
			if args.Len() == 0 {
				return errors.New("expected path to atleast 1 image")
			}

			for i := range args.Len() {
				if err = project.AddImage(args.Get(i)); err != nil {
					return err
				}
			}

			return project.Save()
		},
	}
}
