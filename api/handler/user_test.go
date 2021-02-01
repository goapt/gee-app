package handler

import (
	"net/http"
	"testing"

	"github.com/goapt/dbunit"
	"github.com/goapt/gee"
	"github.com/goapt/redis"
	"github.com/goapt/test"
	"github.com/stretchr/testify/require"

	"app/api/middleware"
	"app/api/session"
	"app/provider/user/model"
	"app/provider/user/repo"
	"app/testutil"
)

const (
	testAccessToken = "testtoken123123"
)

func TestUser_Get(t *testing.T) {
	dbunit.New(t, func(d *dbunit.DBUnit) {
		db := d.NewDatabase(testutil.Schema())
		rds := initSession(t)

		repoUser := repo.NewUsers(db, rds)
		handler := NewUser(repoUser, rds)

		req := test.NewRequest("/api/user/get", gee.HandlerFunc(middleware.NewSession(rds)), handler.Get)
		req.Header.Add("Access-Token", testAccessToken)
		resp, err := req.JSON(map[string]interface{}{"user_name": "test"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.Code)
		require.Equal(t, `{"id":1,"user_name":"test","password":"123123","status":1,"created_at":"2019-03-11 22:12:16","updated_at":"2019-03-11 22:12:16"}`, resp.GetBodyString())
	})
}

func initSession(t *testing.T) *redis.Redis {
	rds := test.NewRedis()
	// 构造session数据
	sess := session.New(rds)
	user := &model.Users{
		Id: 1,
	}
	sess.User = user
	err := sess.Save(testAccessToken)
	require.NoError(t, err)
	return rds
}

func TestUser_Login(t *testing.T) {
	dbunit.New(t, func(d *dbunit.DBUnit) {
		db := d.NewDatabase(testutil.Schema())
		rds := test.NewRedis()
		repoUser := repo.NewUsers(db, rds)
		handler := NewUser(repoUser, rds)

		req := test.NewRequest("/login", handler.Login)
		resp, err := req.JSON(map[string]interface{}{"user_name": "test", "password": "123123"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.Code)
		require.Equal(t, 32, len(resp.GetJsonBody("access_token").String()))
	})
}
