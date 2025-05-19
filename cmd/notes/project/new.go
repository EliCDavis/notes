package project

import (
	"errors"
	"fmt"
	"os"
	"strconv"

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
			&cli.StringFlag{
				Name:  "logs",
				Usage: "path to logs",
				Value: "logs",
			},
			&cli.StringFlag{
				Name:  "tasks",
				Usage: "path to tasks",
				Value: "tasks",
			},
			&cli.StringFlag{
				Name:  "meetings",
				Usage: "path to meetings",
				Value: "meetings",
			},
			&cli.StringFlag{
				Name:  "topics",
				Usage: "path to topics",
				Value: "topics",
			},
			&cli.StringFlag{
				Name:  "images",
				Usage: "path to images",
				Value: "images",
			},
			&cli.StringFlag{
				Name:  "builds",
				Usage: "path to compiled project reports",
				Value: "builds",
			},
			&cli.StringFlag{
				Name:  "mode",
				Usage: "file mode for created directories",
				Value: "0755",
			},
		},
		Action: func(ctx *cli.Context) error {
			args := ctx.Args()
			if args.Len() != 1 {
				return errors.New("expected project name")
			}

			modeStr := ctx.String("mode")
			modeVal, err := strconv.ParseUint(modeStr, 8, 32)
			if err != nil {
				return fmt.Errorf("invalid file mode: %w", err)
			}
			mode := os.FileMode(modeVal)

			newProject := notes.Project{
				Name:         args.First(),
				LogsPath:     ctx.String("logs"),
				TasksPath:    ctx.String("tasks"),
				TopicsPath:   ctx.String("topics"),
				MeetingsPath: ctx.String("meetings"),
				ImagesPath:   ctx.String("images"),
				BuildsPath:   ctx.String("builds"),
			}

			return newProject.SetupFS(ctx.String("path"), mode)
		},
	}
}
