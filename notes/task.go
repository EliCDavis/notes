package notes

import (
	"fmt"
	"os"
	"path"
	"time"
)

const taskFileName = "Description.md"

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
	Name    string
	Path    string
	Created time.Time
	History []*TaskStatusChange `json:"history"`
}

func (t Task) initiailzeMarkdown(parentFolder string) error {
	folder := path.Join(parentFolder, t.Path)
	err := os.MkdirAll(folder, os.ModeDir)
	if err != nil {
		return fmt.Errorf("unable to create tasks's directory: %w", err)
	}

	markdownPath := path.Join(folder, taskFileName)
	file, err := os.Create(markdownPath)
	if err != nil {
		return fmt.Errorf("unable to create log's markdown file: %w", err)
	}
	defer file.Close()

	openURL(markdownPath)

	_, err = fmt.Fprintf(file, "<!-- Created: %s --> \n", t.Created)
	return err
}
