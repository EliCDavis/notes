package flags

import (
	"github.com/EliCDavis/notes/notes"
	"github.com/urfave/cli/v2"
)

var Project = &cli.StringFlag{
	Name:  "project",
	Usage: "Path to project.json",
	Value: "./project.json",
}

func LoadProject(ctx *cli.Context) (*notes.Project, error) {
	projetPath := ctx.String("project")
	return notes.LoadProject(projetPath)
}
