package cache

import (
	"testing"
	"time"

	"github.com/goapt/test"
	"github.com/stretchr/testify/assert"

	"app/provider/user/model"
)

func TestUsers_GetUser(t *testing.T) {
	rds := test.NewRedis()
	u := NewUsers(rds)

	m := &model.Users{
		Id:       1,
		UserName: "test",
		Password: "123123",
		Status:   1,
		CreateAt: time.Time{},
		UpdateAt: time.Time{},
	}
	err := u.SetUser(m)
	assert.NoError(t, err)
	m2 := &model.Users{}
	err = u.GetUser(1, m2)
	assert.NoError(t, err)
	assert.Equal(t, m.UserName, m2.UserName)
}
