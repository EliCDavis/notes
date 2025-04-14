package log

import (
	"github.com/urfave/cli/v2"
)

func newCommand() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "Creates a new log",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			project, err := loadProject(ctx)
			if err != nil {
				return err
			}

			if err = project.NewLog(nil); err != nil {
				return err
			}

			return project.Save()
		},
	}
}
