package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique;column:username"` //设计数据库的一个操作gorm只能出现一次
	Password string `gorm:"column:password"`
	Follow   int64
	Follower int64
	Avatar   string
}

func (user *User) CheckPassword(password string) bool {
	if password == user.Password {
		return true
	}
	return false
}

func (user *User) SetPassword(password string) error {
	user.Password = password
	return nil
}
