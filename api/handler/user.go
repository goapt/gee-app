package handler

import (
	"fmt"
	"time"

	"app/api/constant"
	"github.com/goapt/gee"
	"github.com/goapt/golib/hashing"
	"github.com/goapt/redis"

	"app/api/response"
	"app/api/session"
	"app/provider/user"
	"app/provider/user/model"
	"app/provider/user/repo"
)

type User struct {
	repoUser *repo.Users
	rds      *redis.Redis
}

func NewUser(repoUser *repo.Users, rds user.Redis) *User {
	return &User{repoUser: repoUser, rds: rds}
}

func (u *User) Login(c *gee.Context) gee.Response {
	p := &struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(p); err != nil {
		return response.ParamError(c, err.Error())
	}

	um := &model.Users{
		UserName: p.UserName,
	}

	err := u.repoUser.Find(um)
	if err != nil {
		return response.WithSystemError(c, constant.ErrUserNotExists, err)
	}

	if um.Password != p.Password {
		return response.ParamError(c, "password error")
	}

	token := hashing.Md5(fmt.Sprintf("%d-%d", um.Id, time.Now().UnixNano()))

	sess := session.New(u.rds)
	sess.User = um
	err = sess.Save(token)

	if err != nil {
		return response.Error(c, constant.ErrLoginFail, "登录失败")
	}

	return c.JSON(gee.H{
		"access_token": token,
	})
}

func (u *User) Get(c *gee.Context) gee.Response {
	p := &struct {
		UserName string `json:"user_name" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(p); err != nil {
		return response.ParamError(c, err.Error())
	}

	um := &model.Users{
		UserName: p.UserName,
	}

	err := u.repoUser.Find(um)
	if err != nil {
		return response.WithSystemError(c, constant.ErrUserNotExists, err)
	}

	return c.JSON(um)
}
