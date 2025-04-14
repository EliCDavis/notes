package project

import "github.com/urfave/cli/v2"

func Command() *cli.Command {
	return &cli.Command{
		Name:  "project",
		Usage: "Project management functionality",
		Subcommands: []*cli.Command{
			newCommand(),
		},
	}
}
