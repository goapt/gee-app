package middleware

import (
	"github.com/google/wire"
)

type Middleware struct {
	AccessLog
	Recover
	Cors
	Limiter
	Session
}

var ProviderSet = wire.NewSet(NewCors, NewLimiter, NewRecover, NewSession, NewAccessLog, wire.Struct(new(Middleware), "*"))
