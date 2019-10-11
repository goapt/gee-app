package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/goapt/gee"

	"app/config"
	"app/util"
)

var TokenMiddleware = func() gin.HandlerFunc {
	var fn gee.HandlerFunc = func(c *gee.Context) gee.Response {
		token := c.GetToken()

		if token == "" {
			return c.Fail(40001, "请先登录")
		}

		userId, err := util.AesDecrypt(config.App.TokenSecret, token)

		if err != nil {
			return c.Fail(40003, "Access Token错误"+err.Error())
		}

		c.Set("user_id", userId)

		c.Next()
		return nil
	}
	return gee.Middleware(fn)
}
