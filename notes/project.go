package notes

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const DefaultProjectFile = "project.json"

func LoadProject(projectPath string) (*Project, error) {
	project, err := LoadJSON[Project](projectPath)
	if err != nil {
		return nil, err
	}
	project.loadedPath = projectPath
	return &project, nil
}

type Project struct {
	Name     string `json:"name"`
	LogsPath string `json:"logPath"`
	Logs     []Log  `json:"logs"`

	// The path to this project on disk
	loadedPath string
}

func (p Project) SetupFS(parentFolder string) error {
	projectFolderName := removeNonAlphanumeric(p.Name)
	folder := filepath.Join(parentFolder, projectFolderName)
	err := os.Mkdir(folder, os.ModeDir)
	if err != nil {
		return fmt.Errorf("unable to create project folder %s: %w", folder, err)
	}

	return SaveJSON(filepath.Join(folder, DefaultProjectFile), p)
}

func (p *Project) NewLog(tags []string) error {
	t := time.Now()
	log := Log{
		Time: t,
		Path: t.Format("2006-01-02"),
		Tags: sanitizeTags(tags),
	}

	err := log.initiailzeMarkdown(filepath.Join(filepath.Dir(p.loadedPath), p.LogsPath))
	if err != nil {
		return fmt.Errorf("error creating log for project %s: %w", p.Name, err)
	}

	p.Logs = append(p.Logs, log)
	return nil
}

func (p *Project) Save() error {
	return SaveJSON(p.loadedPath, p)
}
