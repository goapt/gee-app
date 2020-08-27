package connect

import (
	"github.com/goapt/redis"
	"github.com/google/wire"
	"github.com/ilibs/gosql/v2"

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

var ProviderSet = wire.NewSet(NewUserDB, NewUserRedis)
