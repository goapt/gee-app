package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/goapt/gee"

	"app/cmd"
	"app/config"
)

func main() {
	config.Bootstrap()
	//server setup
	cliServ := gee.NewCliServer()
	cliServ.Run(cmd.Commands())
}
