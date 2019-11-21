package session

import (
	"github.com/goapt/gee"
	"github.com/ilibs/gosql/v2"

	"app/model"
)

// 定义session handle
type Handler func(u *model.Users, c *gee.Context) gee.Response

func (h Handler) Handle(c *gee.Context) gee.Response {
	user := &model.Users{}
	err := gosql.Model(user).Where("id = ?", c.GetString("user_id")).Get()
	if err != nil {
		return c.Fail(40002, err)
	}
	return h(user, c)
}
