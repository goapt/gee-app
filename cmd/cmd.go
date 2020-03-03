package cmd

import (
	"github.com/urfave/cli"
)

var cmds cli.Commands

func register(cmd cli.Command) {
	cmds = append(cmds, cmd)
}

func Commands() cli.Commands {
	return cmds
}
