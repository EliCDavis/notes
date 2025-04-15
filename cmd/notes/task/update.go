package task

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/EliCDavis/notes/notes"
	"github.com/urfave/cli/v2"
)

func updateCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "update a task",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "task",
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
						return errors.New("Expected name")
					}

					project.Tasks[ctx.Int("task")-1].Name = args.First()
					return project.Save()
				},
			},

			{
				Name:  "status",
				Usage: "Update the status of the task",
				Action: func(ctx *cli.Context) error {
					project, err := flags.LoadProject(ctx)
					if err != nil {
						return err
					}

					args := ctx.Args()
					if args.Len() != 1 {
						return errors.New("Expected name")
					}

					var newStatus notes.TaskStatus
					switch strings.ToLower(strings.TrimSpace(args.First())) {
					case "start":
						newStatus = notes.TaskStatus_Started

					case "complete":
						newStatus = notes.TaskStatus_Completed

					case "stop":
						newStatus = notes.TaskStatus_Stopped

					case "abandon":
						newStatus = notes.TaskStatus_Abandoned

					default:
						return fmt.Errorf("unrecognized status '%s'", args.First())
					}

					task := project.Tasks[ctx.Int("task")-1]
					task.History = append(task.History, &notes.TaskStatusChange{
						Status: newStatus,
						Time:   time.Now(),
					})
					return project.Save()
				},
			},
		},
	}
}
