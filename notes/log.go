package notes

import (
	"fmt"
	"os"
	"path"
	"time"
)

const logFileName = "README.md"

type Log struct {
	Time time.Time `json:"time"`
	Path string    `json:"path"`
	Tags []string  `json:"tags"`
}

func (l Log) initiailzeMarkdown(parentFolder string) error {
	folder := path.Join(parentFolder, l.Path)
	err := os.MkdirAll(folder, os.ModeDir)
	if err != nil {
		return fmt.Errorf("unable to create log's directory: %w", err)
	}

	file, err := os.Create(path.Join(folder, logFileName))
	if err != nil {
		return fmt.Errorf("unable to create log's markdown file: %w", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "# %s\n\n", l.Time.Format("2006-01-02"))
	return err
}
