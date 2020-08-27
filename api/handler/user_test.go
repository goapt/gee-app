package handler

import (
	"net/http"
	"testing"

	"github.com/goapt/dbunit"
	"github.com/goapt/test"
	"github.com/stretchr/testify/require"

	"app/provider/user/repo"
	"app/testutil"
)

func TestUser_Get(t *testing.T) {
	dbunit.New(t, func(d *dbunit.DBUnit) {
		db := d.NewDatabase(testutil.Schema())
		rds := test.NewRedis()

		repoUser := repo.NewUsers(db, rds)
		handler := NewUser(repoUser, rds)

		req := test.NewRequest("/api/user/get", handler.Get)
		resp, err := req.JSON(map[string]interface{}{"user_name": "test"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.Code)
		require.Equal(t, `{"id":1,"user_name":"test","password":"123123","status":1,"created_at":"2019-03-11 22:12:16","updated_at":"2019-03-11 22:12:16"}`, resp.GetBodyString())

	})
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
		require.Equal(t, `{"access_token":"4xgCqZpNHGyEwSHshM6ocg=="}`, resp.GetBodyString())
	})
}
