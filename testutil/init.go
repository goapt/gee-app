package testutil

import (
	"github.com/goapt/golib/debug"
	"github.com/goapt/logger"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	debug.Open("on", "")
	logger.Setting(func(c *logger.Config) {
		c.LogMode = "std"
	})
}
