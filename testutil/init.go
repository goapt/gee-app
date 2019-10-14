package testutil

import (
	"github.com/goapt/golib/debug"
	"github.com/goapt/logger"

	"app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilibs/gosql/v2"
)

func init() {
	//db connect
	_ = gosql.Connect(config.App.DB)

	debug.Open("on", "")
	logger.Setting(func(c *logger.Config) {
		c.LogMode = "std"
	})
}
