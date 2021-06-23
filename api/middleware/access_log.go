package middleware

import (
	"app/connect"
	"github.com/goapt/gee"
	"github.com/goapt/logger"
)

type AccessLog gee.HandlerFunc

func NewAccessLog(log connect.AccessLogger) AccessLog {
	return func(c *gee.Context) gee.Response {
		entry := logger.NewAccessLog(c.Request)
		entry.StartTime = c.StartTime
		entry.Response = c.ResponseWriter()
		entry.LogInfo = c.LogInfo
		c.Next()
		if data := entry.Get(); data != nil {
			log.WithFields(entry.Get()).Info("access_log")
		}
		return nil
	}
}
