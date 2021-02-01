package middleware

import (
	"net/http"

	"github.com/goapt/gee"
	"github.com/goapt/golib/debug"
	"github.com/goapt/logger"
)

type Recover gee.HandlerFunc

func NewRecover() Recover {
	return func(c *gee.Context) gee.Response {
		defer func(c *gee.Context) {
			if rec := recover(); rec != nil {
				logger.Data(map[string]interface{}{
					"error": rec,
					"path":  c.Request.URL.Path,
					"stack": string(debug.Stack(-1)),
				}).Error("[golang painc]", c.Request.URL.Path)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}(c)
		c.Next()
		return nil
	}
}
