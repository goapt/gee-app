package middleware

import (
	"github.com/goapt/redis"
	"github.com/google/wire"

	"app/provider/user"
)

type Middleware struct {
	userRedis *redis.Redis
}

func NewMiddleware(userRedis user.Redis) *Middleware {
	return &Middleware{userRedis: userRedis}
}

var ProviderSet = wire.NewSet(NewMiddleware)
