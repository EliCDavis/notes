package notes

import (
	"fmt"
	"os"
	"path"
	"time"

	_ "embed"
)

const meetingFileName = "README.md"

//go:embed default_meeting.md
var defaultMeeting []byte

type Meeting struct {
	Path    string
	Created time.Time
}

func (l Meeting) initiailzeMarkdown(parentFolder string) error {
	folder := path.Join(parentFolder, l.Path)
	err := os.MkdirAll(folder, os.ModeDir)
	if err != nil {
		return fmt.Errorf("unable to create meeting's directory: %w", err)
	}

	meetingPath := path.Join(folder, meetingFileName)
	file, err := os.Create(meetingPath)
	if err != nil {
		return fmt.Errorf("unable to create meeting's markdown file: %w", err)
	}
	defer file.Close()

	openURL(meetingPath)

	_, err = file.Write(defaultMeeting)
	return err
}
