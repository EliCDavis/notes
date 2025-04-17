package notes

import (
	"fmt"
	"os"
	"path"
)

const topicFileName = "README.md"

type Topic struct {
	Entry
	Name string
}

func (l Topic) initiailzeMarkdown(parentFolder string) error {
	folder := path.Join(parentFolder, l.Path)
	err := os.MkdirAll(folder, os.ModeDir)
	if err != nil {
		return fmt.Errorf("unable to create topic's directory: %w", err)
	}

	topicPath := path.Join(folder, topicFileName)
	file, err := os.Create(topicPath)
	if err != nil {
		return fmt.Errorf("unable to create topic's markdown file: %w", err)
	}
	defer file.Close()

	openURL(topicPath)
	_, err = fmt.Fprintf(file, "<!-- Created: %s --> \n", l.Created)
	return err
}
