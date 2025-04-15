package meeting

import (
	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func newCommand() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "Creates a new meeting",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			if err = project.NewMeeting(); err != nil {
				return err
			}

			return project.Save()
		},
	}
}
