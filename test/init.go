package test

import (
	"github.com/goapt/golib/debug"
	"github.com/goapt/logger"

	_ "github.com/go-sql-driver/mysql"

	"app/config"
)

func init() {
	config.Bootstrap()

	debug.Open("on", "")
	logger.Setting(func(c *logger.Config) {
		c.LogMode = "std"
	})
}
