package handler

import (
	"strconv"

	"github.com/goapt/gee"
	"github.com/goapt/redis"

	"app/api/response"
	"app/api/session"
	"app/config"
	"app/pkg/cryptutil"
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
		return response.Error(c, err)
	}

	if um.Password != p.Password {
		return response.ParamError(c, "password error")
	}

	token, err := cryptutil.AesEncrypt(config.App.TokenSecret, strconv.Itoa(um.Id))

	if err != nil {
		return response.Error(c, err)
	}

	sess := session.New(u.rds)
	sess.User = um
	err = sess.Save(strconv.Itoa(um.Id))

	if err != nil {
		return response.Error(c, err)
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

	user := &model.Users{
		UserName: p.UserName,
	}

	err := u.repoUser.Find(user)
	if err != nil {
		return response.Error(c, err)
	}

	return c.JSON(user)
}
