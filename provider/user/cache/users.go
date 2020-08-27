package cache

import "github.com/goapt/redis"

type Users struct {
	*redis.Redis
}

func NewUsers(rds *redis.Redis) *Users {
	return &Users{
		rds,
	}
}
