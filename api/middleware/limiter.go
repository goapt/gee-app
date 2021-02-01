package middleware

import (
	"github.com/goapt/gee"
	"golang.org/x/time/rate"

	"app/api/response"
)

type Limiter gee.HandlerFunc

// NewLimiter 限流中间件
func NewLimiter(limiter *rate.Limiter) Limiter {
	return func(c *gee.Context) gee.Response {
		if limiter.Allow() {
			c.Next()
			return nil
		}
		c.Abort()
		return response.Error(c, response.ErrRateLimited)
	}
}
