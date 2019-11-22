package cmd

import (
	"github.com/urfave/cli"
)

var cmds cli.Commands

func registerCommand(cmd cli.Command) {
	cmds = append(cmds, cmd)
}

func Boostrap() cli.Commands {
	registerCommand(httpCmd)
	registerCommand(testCmd)
	return cmds
}
