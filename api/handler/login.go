package handler

import (
	"strconv"

	"app/api/session"
	"app/config"
	"app/model"
	"app/pkg/cryptutil"

	"github.com/gin-gonic/gin"
	"github.com/goapt/gee"
	"github.com/ilibs/gosql/v2"
)

// 登录代码
var LoginHandle gee.HandlerFunc = func(c *gee.Context) gee.Response {
	p := &struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(p); err != nil {
		return c.Fail(201, err)
	}

	user := &model.Users{
		UserName: p.UserName,
	}

	err := gosql.Model(user).Get()

	if err != nil {
		return c.Fail(201, err)
	}

	if user.Password != p.Password {
		return c.Fail(201, "密码错误")
	}

	token, err := cryptutil.AesEncrypt(config.App.TokenSecret, strconv.Itoa(user.Id))

	if err != nil {
		return c.Fail(202, err)
	}

	// 构造session数据
	sess := session.New()
	sess.User = user
	err = sess.Save(strconv.Itoa(user.Id))

	if err != nil {
		return c.Fail(201, "session save error")
	}

	return c.Success(gin.H{
		"access_token": token,
	})
}
