package middleware

import (
	"testing"
	"time"

	"github.com/goapt/gee"
	"github.com/goapt/test"
	"github.com/stretchr/testify/assert"

	"app/api/session"
	"app/provider/user/model"
)

func TestMiddleware_Session(t *testing.T) {
	var testHandler gee.HandlerFunc = func(c *gee.Context) gee.Response {
		u := session.Init(c).User
		return c.JSON(u)
	}

	rds := test.NewRedis()

	sess := session.New(rds)
	sess.User = &model.Users{
		Id:       1,
		UserName: "test",
		Password: "123123",
		Status:   1,
		CreateAt: time.Time{},
		UpdateAt: time.Time{},
	}
	accessToken := "4xgCqZpNHGyEwSHshM6ocg=="
	err := sess.Save("1")
	assert.NoError(t, err)

	mid := NewMiddleware(rds)

	req := test.NewRequest("/dummy/impl", mid.Session(), testHandler)
	req.Header.Add("Access-Token", accessToken)
	resp, err := req.JSON(nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, `test`, resp.GetJsonBody("user_name").String())
}
