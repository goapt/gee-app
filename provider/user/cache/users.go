package cache

import (
	"fmt"

	"github.com/goapt/redis"

	"app/provider/user/model"
)

type Users struct {
	*redis.Redis
}

func NewUsers(rds *redis.Redis) *Users {
	return &Users{
		rds,
	}
}

func (u *Users) GetUser(userId int, model *model.Users) error {
	key := fmt.Sprintf("user:%d", userId)
	return u.HGetAll(key, model)
}

func (u *Users) SetUser(model *model.Users) error {
	key := fmt.Sprintf("user:%d", model.Id)
	return u.HMSet(key, model)
}
