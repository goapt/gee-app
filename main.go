package main

import (
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goapt/gee"
	"github.com/goapt/logger"

	"app/config"
	"app/connect"
)

func main() {
	connect.Connect(config.App)

	// logger init
	logger.Setting(func(c *logger.Config) {
		c.LogName = "app"
		c.LogMode = config.App.Log.LogMode
		c.LogPath = filepath.Join(config.App.Log.LogPath, config.App.AppName)
		c.LogLevel = config.App.Log.LogLevel
		c.LogMaxFiles = config.App.Log.LogMaxFiles
		c.LogSentryDSN = config.App.Log.LogSentryDSN
		c.LogSentryType = "go." + config.App.AppName
		c.LogDetail = config.App.Log.LogDetail
	})

	cmds := Initialize()
	// server setup
	gee.NewCliServer().Run(cmds)
}
