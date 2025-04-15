package project

import (
	"bufio"
	"fmt"
	"io"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/urfave/cli/v2"
)

func compileCommand() *cli.Command {
	return &cli.Command{
		Name:  "compile",
		Usage: "compiles project into a single file",
		Flags: []cli.Flag{
			flags.Project,
		},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			var out io.Writer = ctx.App.Writer
			writer := bufio.NewWriter(out)
			err = project.Compile(writer)
			if err != nil {
				return fmt.Errorf("unable to compile project: %w", err)
			}
			return writer.Flush()
		},
	}
}
