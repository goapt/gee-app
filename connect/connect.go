package connect

import (
	"path/filepath"
	"time"

	"github.com/goapt/logger"
	"github.com/goapt/redis"
	"github.com/google/wire"
	"github.com/ilibs/gosql/v2"
	"golang.org/x/time/rate"

	"app/config"
	"app/provider/user"
)

func Connect(c *config.Config) {
	// db connection
	_ = gosql.Connect(c.DB)

	// redis connection
	redis.Connect(c.Redis)
}

func NewUserDB() user.DB {
	return gosql.Use("default")
}

func NewUserRedis() user.Redis {
	return redis.NewRedisWithName("default")
}

func NewRateLimiter() *rate.Limiter {
	return rate.NewLimiter(rate.Every(time.Second*1), 100000) // 根据项目配置，可以不开启
}

type AccessLogger logger.ILogger

func NewAccessLogger() AccessLogger {
	newLog := logger.NewLogger(func(c *logger.Config) {
		c.LogName = "access"
		c.LogMode = config.App.Log.LogMode
		c.LogPath = filepath.Join(config.App.Log.LogPath, config.App.AppName)
		c.LogLevel = config.App.Log.LogLevel
		c.LogMaxFiles = config.App.Log.LogMaxFiles
		c.LogSentryDSN = ""
		c.LogDetail = false
	})
	return newLog
}

var ProviderSet = wire.NewSet(NewUserDB, NewUserRedis, NewRateLimiter, NewAccessLogger)
