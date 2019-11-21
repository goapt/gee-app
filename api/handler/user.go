package handler

import (
	"github.com/goapt/gee"

	"app/api/session"
	"app/model"
)

var UserHandle session.Handler = func(s *model.Users, c *gee.Context) gee.Response {
	return c.Success(s)
}
