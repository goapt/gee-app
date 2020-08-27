package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goapt/gee"
	"github.com/goapt/golib/debug"
	"github.com/goapt/logger"
)

func (m *Middleware) Recover() gee.HandlerFunc {
	return gee.Wrap(func(c *gin.Context) {
		defer func(c *gin.Context) {
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
	})
}
