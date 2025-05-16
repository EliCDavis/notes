package tag

import (
	"errors"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func updateCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "update a tag",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "tag",
				Usage:    "ID of task to update",
				Required: true,
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:      "name",
				Usage:     "Update the name of the task",
				Args:      true,
				ArgsUsage: "[New Name]",
				Action: func(ctx *cli.Context) error {
					project, err := flags.LoadProject(ctx)
					if err != nil {
						return err
					}

					args := ctx.Args()
					if args.Len() != 1 {
						return errors.New("expected name")
					}

					project.Tags[ctx.Int("tag")-1].Name = args.First()
					return project.Save()
				},
			},
		},
	}
}
