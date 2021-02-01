package connect

import (
	"time"

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

var ProviderSet = wire.NewSet(NewUserDB, NewUserRedis, NewRateLimiter)
