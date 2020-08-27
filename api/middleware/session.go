package middleware

import (
	"errors"

	"github.com/goapt/gee"
	"github.com/goapt/logger"

	"app/api/response"
	"app/api/session"
	"app/config"
	"app/pkg/cryptutil"
)

func parseToken(c *gee.Context, secret string) (string, error) {
	token := c.Request.Header.Get("Access-Token")
	if token == "" || len(token) < 15 {
		return "", errors.New("access denied")
	}

	id, err := cryptutil.AesDecrypt(secret, token)

	if err != nil {
		return "", err
	}

	if id == "" {
		return "", errors.New("请先登录")
	}

	return id, nil
}

func (m *Middleware) Session() gee.HandlerFunc {
	return func(c *gee.Context) gee.Response {
		sess := session.New(m.userRedis)
		var (
			id  string
			err error
		)

		if id, err = parseToken(c, config.App.TokenSecret); err != nil {
			c.Abort()
			return response.ParamError(c, err.Error())
		}

		_, err = sess.Get(id)

		if err != nil {
			c.Abort()
			return response.Error(c, response.ErrAuthFailure, "登录超时请重新登录")
		}

		c.Set("__session", sess)
		c.Next()

		if err := sess.Save(id); err != nil {
			logger.Error("save session error", err)
		}

		return nil
	}
}
