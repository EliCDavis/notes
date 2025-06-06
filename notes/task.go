package notes

import (
	"bufio"
	"fmt"
	"time"
)

type TaskStatus string

const (
	TaskStatus_Started   = "Started"
	TaskStatus_Stopped   = "Stopped"
	TaskStatus_Completed = "Completed"
	TaskStatus_Abandoned = "Abandoned"
)

type TaskStatusChange struct {
	Status TaskStatus
	Time   time.Time
	Reason string
}

type Task struct {
	Entry
	Name    string
	History []*TaskStatusChange `json:"history"`
}

func (t Task) DisplayName() string {
	if t.Name == "" {
		return "[Unnamed]"
	}
	return t.Name
}

func (t Task) toProjectMarkdown(writer *bufio.Writer) error {
	for _, item := range t.History {
		fmt.Fprintf(writer, "* %s: %s", item.Status, item.Time)

		if item.Reason != "" {
			fmt.Fprintf(writer, " - %s", item.Reason)
		}
		fmt.Fprint(writer, "\n")
	}

	if len(t.History) > 0 {
		fmt.Fprint(writer, "\n")
	}

	return nil
}
