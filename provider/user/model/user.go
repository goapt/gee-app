package model

import "time"

type Users struct {
	Id       int       `json:"id" db:"id"`                                                   // 用户ID
	UserName string    `json:"user_name" db:"user_name"`                                     // 登录名
	Password string    `json:"password" db:"password"`                                       // 登录密码
	Status   int       `json:"status" db:"status"`                                           // 1正常 0停用，2待激活
	CreateAt time.Time `json:"created_at" db:"created_at" time_format:"2006-01-02 15:04:05"` // 创建时间
	UpdateAt time.Time `json:"updated_at" db:"updated_at" time_format:"2006-01-02 15:04:05"` // 修改时间
}

func (this *Users) TableName() string {
	return "users"
}

func (this *Users) PK() string {
	return "user_id"
}
