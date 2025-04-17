package notes

import (
	"fmt"
	"os"
	"path"
)

const logFileName = "README.md"

type Log Entry

func (l Log) initiailzeMarkdown(parentFolder string) error {
	folder := path.Join(parentFolder, l.Path)
	err := os.MkdirAll(folder, os.ModeDir)
	if err != nil {
		return fmt.Errorf("unable to create log's directory: %w", err)
	}

	logPath := path.Join(folder, logFileName)
	file, err := os.Create(logPath)
	if err != nil {
		return fmt.Errorf("unable to create log's markdown file: %w", err)
	}
	defer file.Close()

	openURL(logPath)

	_, err = fmt.Fprintf(file, "<!-- Created: %s --> \n", l.Created)
	return err
}
