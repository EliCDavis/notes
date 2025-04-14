package log

import (
	"github.com/urfave/cli/v2"
)

func compileCommand() *cli.Command {
	return &cli.Command{
		Name:  "compile",
		Usage: "compiles logs into a single file",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}
}
