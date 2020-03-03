package middleware

import (
	"errors"

	"github.com/goapt/gee"
	"github.com/goapt/logger"

	"app/api/session"
	"app/config"
	"app/pkg/cryptutil"
)

func parseToken(c *gee.Context) (string, error) {
	token := c.Request.Header.Get("Access-Token")
	if token == "" || len(token) < 15 {
		return "", errors.New("access denied")
	}

	id, err := cryptutil.AesDecrypt(config.App.TokenSecret, token)

	if err != nil {
		return "", err
	}

	if id == "" {
		return "", errors.New("请先登录")
	}

	return id, nil
}

var SessionMiddleware = func() gee.HandlerFunc {
	sess := session.New()

	return func(c *gee.Context) gee.Response {
		var (
			id  string
			err error
		)

		if id, err = parseToken(c); err != nil {
			c.Abort()
			return c.Fail(40001, err)
		}

		_, err = sess.Get(id)

		if err != nil {
			c.Abort()
			return c.Fail(40002, err)
		}

		c.Set("__session", sess)
		c.Next()

		if err := sess.Save(id); err != nil {
			logger.Error("save session error", err)
		}

		return nil
	}
}
