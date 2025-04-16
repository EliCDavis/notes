package project

import (
	"bufio"
	"fmt"
	"io"

	"github.com/EliCDavis/notes/cmd/notes/flags"
	"github.com/EliCDavis/notes/notes"
	"github.com/urfave/cli/v2"
)

const useIncludesFlag = "use-includes"
const saveFlag = "save"

func compileCommand() *cli.Command {
	return &cli.Command{
		Name:  "compile",
		Usage: "compiles project into a single file",
		Flags: []cli.Flag{
			flags.Project,
			&cli.BoolFlag{
				Name:  useIncludesFlag,
				Usage: "Whether or not to use the markdown-it-include extension",
			},
			&cli.BoolFlag{
				Name:  saveFlag,
				Usage: "Whether or not to save compiled report",
			},
		},
		Action: func(ctx *cli.Context) error {
			project, err := flags.LoadProject(ctx)
			if err != nil {
				return err
			}

			var out io.Writer = ctx.App.Writer
			writer := bufio.NewWriter(out)
			err = project.Compile(writer, notes.ProjectCompileOptions{
				UseMarkdownItIncludeExtension: ctx.Bool(useIncludesFlag),
				Save:                          ctx.Bool(saveFlag),
			})
			if err != nil {
				return fmt.Errorf("unable to compile project: %w", err)
			}
			return writer.Flush()
		},
	}
}
