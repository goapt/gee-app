package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/goapt/gee"
	"github.com/ilibs/gosql/v2"

	_ "app/cmd"
	"app/config"
)

func main() {
	//db connect
	_ = gosql.Connect(config.App.DB)

	//command server
	cliServ := gee.NewCliServer()
	cliServ.Run()
}
