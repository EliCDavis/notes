package topic

import (
	"errors"
	"fmt"
	"strings"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func newCommand() *cli.Command {
	return &cli.Command{
		Name:      "new",
		Usage:     "Creates a new topic",
		Args:      true,
		ArgsUsage: "[Topic Name]",
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			args := ctx.Args()
			if args.Len() != 1 {
				return errors.New("expected 1 argument for topic name")
			}

			name := strings.TrimSpace(args.First())
			if name == "" {
				return fmt.Errorf("invalid name: %s", args.First())
			}

			if err = project.NewTopic(name); err != nil {
				return err
			}

			return project.Save()
		},
	}
}
