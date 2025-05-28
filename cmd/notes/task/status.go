package task

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/EliCDavis/notes/notes"
	"github.com/urfave/cli/v2"
)

func updateStatus(ctx *cli.Context, status notes.TaskStatus) error {
	project, err := flags.LoadProject(ctx)
	if err != nil {
		return err
	}

	args := ctx.Args()
	if args.Len() == 0 {
		return errors.New("expected task id")
	}

	taskId, err := strconv.Atoi(args.Get(0))
	if err != nil {
		return fmt.Errorf("unable to parse task ID %q: %w", args.Get(0), err)
	}

	reason := ""
	if args.Len() > 1 {
		reason = args.Get(1)
	}

	task := project.Tasks[taskId-1]
	task.History = append(task.History, &notes.TaskStatusChange{
		Status: status,
		Time:   time.Now(),
		Reason: reason,
	})
	return project.Save()
}

func startCommand() *cli.Command {
	return &cli.Command{
		Name:      "start",
		Usage:     "start a task",
		Args:      true,
		ArgsUsage: "[Task ID] [reason]",
		Action: func(ctx *cli.Context) error {
			return updateStatus(ctx, notes.TaskStatus_Started)
		},
	}
}

func stopCommand() *cli.Command {
	return &cli.Command{
		Name:      "stop",
		Usage:     "stop a task",
		Args:      true,
		ArgsUsage: "[Task ID] [reason]",
		Action: func(ctx *cli.Context) error {
			return updateStatus(ctx, notes.TaskStatus_Stopped)
		},
	}
}
func completeCommand() *cli.Command {
	return &cli.Command{
		Name:      "complete",
		Usage:     "complete a task",
		Args:      true,
		ArgsUsage: "[Task ID] [reason]",
		Action: func(ctx *cli.Context) error {
			return updateStatus(ctx, notes.TaskStatus_Completed)
		},
	}
}
func abandonCommand() *cli.Command {
	return &cli.Command{
		Name:      "abandon",
		Usage:     "abandon a task",
		Args:      true,
		ArgsUsage: "[Task ID] [reason]",
		Action: func(ctx *cli.Context) error {
			return updateStatus(ctx, notes.TaskStatus_Abandoned)
		},
	}
}
