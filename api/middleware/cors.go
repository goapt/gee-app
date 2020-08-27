package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/goapt/gee"
)

// Cors
func (m *Middleware) Cors() gee.HandlerFunc {
	return gee.Wrap(cors.New(cors.Config{
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
	}))
}
