package main

import (
	"log"
	"os"

	"github.com/EliCDavis/notes/cmd/notes/images"
	plog "github.com/EliCDavis/notes/cmd/notes/log"
	"github.com/EliCDavis/notes/cmd/notes/meeting"
	"github.com/EliCDavis/notes/cmd/notes/project"
	"github.com/EliCDavis/notes/cmd/notes/tag"
	"github.com/EliCDavis/notes/cmd/notes/task"
	"github.com/EliCDavis/notes/cmd/notes/topic"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "notes",
		Usage: "Manage a collection of notes",
		Authors: []*cli.Author{
			{
				Name: "Eli C Davis",
			},
		},
		Commands: []*cli.Command{
			project.Command(),
			plog.Command(),
			task.Command(),
			meeting.Command(),
			topic.Command(),
			images.Command(),
			tag.Command(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
