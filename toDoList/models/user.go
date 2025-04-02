package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 建立用户表
type User struct {
	gorm.Model
	Username       string `gorm:"unique"` // 用户名唯一
	PasswordDigest string //存储的是密文，加密后的密码
}

// 对密码进行加密
func (u *User) SetPassWord(passWord string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passWord), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// 验证密码
func (u *User) CheckPassWord(passWord string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(passWord))
	return err == nil
}
