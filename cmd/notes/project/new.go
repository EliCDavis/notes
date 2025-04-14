package project

import (
	"errors"

	"github.com/EliCDavis/notes/notes"
	"github.com/urfave/cli/v2"
)

func newCommand() *cli.Command {
	return &cli.Command{
		Name:      "new",
		Usage:     "Creates a new project",
		Args:      true,
		ArgsUsage: "[project name]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "path to place project",
				Value: "./",
			},
		},
		Action: func(ctx *cli.Context) error {
			args := ctx.Args()
			if args.Len() != 1 {
				return errors.New("expected project name")
			}

			newProject := notes.Project{
				Name: args.First(),
			}

			return newProject.SetupFS(ctx.String("path"))
		},
	}
}
