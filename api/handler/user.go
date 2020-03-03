package handler

import (
	"github.com/goapt/gee"

	"app/api/session"
)

var UserHandle gee.HandlerFunc = func(c *gee.Context) gee.Response {
	sess, err := session.Init(c)
	if err != nil {
		return c.Fail(201, "session init error")
	}
	return c.Success(sess.User)
}
