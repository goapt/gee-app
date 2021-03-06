package middleware

import (
	"app/api/constant"
	"github.com/goapt/gee"
	"github.com/goapt/logger"

	"app/api/response"
	"app/api/session"
	"app/provider/user"
)

type Session gee.HandlerFunc

func NewSession(rds user.Redis) Session {
	return func(c *gee.Context) gee.Response {
		sess := session.New(rds)
		token := c.Request.Header.Get("Access-Token")
		if token == "" {
			c.Abort()
			return response.Error(c, constant.ErrAccessForbidden, "非法访问")
		}

		_, err := sess.Get(token)

		if err != nil {
			c.Abort()
			return response.SystemError(c, "登录超时请重新登录")
		}

		c.Set("__session", sess)
		c.Next()

		if err := sess.Save(token); err != nil {
			logger.Error("save session error", err)
		}

		return nil
	}
}
