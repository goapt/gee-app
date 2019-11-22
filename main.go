package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/goapt/gee"

	"app/cmd"
	"app/config"
)

func main() {
	cmds := cmd.Boostrap()
	config.Bootstrap()
	//server setup
	cliServ := gee.NewCliServer()
	cliServ.Run(cmds)
}
