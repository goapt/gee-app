package repo

import (
	"app/provider/user"
	"app/provider/user/cache"
	"app/provider/user/model"
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

func (u *Users) GetUser(userId int) (*model.Users, error) {
	m := new(model.Users)
	err := u.userRedis.GetUser(userId, m)
	if err != nil || m.Id == 0 {
		err := u.db.Model(m).Where("id = ?", userId).Get()
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
