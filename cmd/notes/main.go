package main

import (
	"log"
	"os"

	plog "github.com/EliCDavis/notes/cmd/notes/log"
	"github.com/EliCDavis/notes/cmd/notes/project"
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
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
