package tag

import (
	"errors"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func newCommand() *cli.Command {
	return &cli.Command{
		Name:      "new",
		Usage:     "Creates tags for the project",
		Args:      true,
		ArgsUsage: "[name of tags]",
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			args := ctx.Args()
			if args.Len() == 0 {
				return errors.New("expected atleast 1 tag name")
			}

			for i := range args.Len() {
				project.AddTag(args.Get(i))
			}

			return project.Save()
		},
	}
}
