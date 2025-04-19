package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Password     string //存储的的是加密后的密码
	Age          int    `json:"age"`
	Sex          string `json:"sex"`
	BlogCount    int    `json:"blogCount" gorm:"default:0"`
	CommentCount int    `json:"commentCount" gorm:"default:0"`
}

// EncryptPassword 注册时要对输入密码进行加密
func (user *User) EncryptPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 登陆时要验证密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
