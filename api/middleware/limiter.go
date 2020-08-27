package middleware

import (
	"time"

	"github.com/goapt/gee"
	"golang.org/x/time/rate"

	"app/api/response"
)

// LimiterMiddleware 限流中间件
func (m *Middleware) Limiter(maxBurstSize int) gee.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Second*1), maxBurstSize)

	return func(c *gee.Context) gee.Response {
		if limiter.Allow() {
			c.Next()
			return nil
		}
		c.Abort()
		return response.Error(c, response.ErrRateLimited)
	}
}
