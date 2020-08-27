package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/goapt/gee"
	"github.com/goapt/redis"

	"app/config"
	"app/provider/user/model"
)

// 定义session
type Session struct {
	client *redis.Redis
	User   *model.Users
}

func New(rds *redis.Redis) *Session {
	return &Session{
		client: rds,
	}
}

func (*Session) Expire() time.Duration {
	return time.Hour * 64
}

func (*Session) Prefix() string {
	return config.App.AppName
}

func (s *Session) Save(key string) error {
	sess, _ := json.Marshal(s)
	return s.client.SetEX(s.Prefix()+":"+key, string(sess), s.Expire())
}

func (s *Session) Get(key string) (*Session, error) {
	// 获取 session 信息
	val, err := s.client.Get(s.Prefix() + ":" + key)

	if err != nil {
		return nil, err
	}

	// 查询出来之后，还需要再判断ttl是否>0 否则也是过期了
	if s.client.TTL(s.Prefix()+":"+key) <= 0 {
		return nil, errors.New("session已过期")
	}

	if err := json.Unmarshal([]byte(val), s); err != nil {
		return nil, errors.New("Unmarshal error:" + err.Error() + fmt.Sprintf("%+v", s))
	}

	return s, nil
}

func (s *Session) IsLogin() bool {
	return s.User != nil && s.User.Id != 0
}

func Init(c *gee.Context) *Session {
	sess, has := c.Get("__session")
	if !has {
		return nil
	}

	return sess.(*Session)
}
