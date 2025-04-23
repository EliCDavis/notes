package log

import (
	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func openCommand() *cli.Command {
	return &cli.Command{
		Name:  "open",
		Usage: "Opens the current log for the day, or creates one if it hasn't been created yet",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			if err = project.OpenLog(); err != nil {
				return err
			}

			return project.Save()
		},
	}
}
