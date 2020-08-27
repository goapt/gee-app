package repo

import (
	"app/provider/user"
	"app/provider/user/cache"
)

type Users struct {
	Base
	userRedis *cache.Users
}

// Users
func NewUsers(db user.DB, rds user.Redis) *Users {
	return &Users{
		Base:      Base{db: db},
		userRedis: cache.NewUsers(rds),
	}
}
