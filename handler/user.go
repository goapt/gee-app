package handler

import (
	"github.com/goapt/gee"

	"app/model"
	"app/session"
)

var UserHandle session.Handler = func(s *model.Users, c *gee.Context) gee.Response {
	return c.Success(s)
}