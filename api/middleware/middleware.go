package middleware

import (
	"github.com/google/wire"
)

type Middleware struct {
	Recover
	Cors
	Limiter
	Session
}

var ProviderSet = wire.NewSet(NewCors, NewLimiter, NewRecover, NewSession, wire.Struct(new(Middleware), "*"))
