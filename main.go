package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/goapt/gee"

	"app/config"
	"app/connect"
)

func main() {
	connect.Connect(config.App)
	cmds := Initialize()

	// server setup
	gee.NewCliServer().Run(cmds)
}
