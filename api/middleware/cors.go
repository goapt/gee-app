package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/goapt/gee"
)

type Cors gee.HandlerFunc

func NewCors() Cors {
	return func(c *gee.Context) gee.Response {
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"*"},
			AllowHeaders: []string{
				"Origin",
				"Content-Length",
				"Content-Type",
				"Access-Token",
				"Access-Control-Allow-Origin",
				"x-requested-with",
			},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})(c.Context)

		return nil
	}
}
