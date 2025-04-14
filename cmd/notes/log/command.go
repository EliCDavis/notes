package log

import (
	"github.com/EliCDavis/notes/notes"
	"github.com/urfave/cli/v2"
)

func loadProject(ctx *cli.Context) (*notes.Project, error) {
	projetPath := ctx.String("project")
	return notes.LoadProject(projetPath)
}

func Command() *cli.Command {
	return &cli.Command{
		Name:  "log",
		Usage: "Create and edit logs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "project",
				Usage: "Path to project.json",
				Value: "./project.json",
			},
		},
		Subcommands: []*cli.Command{
			newCommand(),
			compileCommand(),
		},
	}
}
