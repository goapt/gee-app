package cmd

import (
	"reflect"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

func NewCommands(cmd *Commands) cli.Commands {
	var cmds cli.Commands
	v := reflect.Indirect(reflect.ValueOf(cmd))
	ct := reflect.TypeOf(&cli.Command{})
	for i := 0; i < v.NumField(); i++ {
		cmds = append(cmds, v.Field(i).Convert(ct).Interface().(*cli.Command))
	}

	return cmds
}

type Commands struct {
	HttpCmd HttpCmd
}

var ProviderSet = wire.NewSet(NewHttp, NewCommands, wire.Struct(new(Commands), "*"))
