package serializer

import "toDoList/models"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	UserName string `json:"user_name" form:"user_name" example:"FanOne"`
	Status   string `json:"status" form:"status"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

// 序列化用户
func BuildUser(user models.User) User {
	return User{
		ID:       user.ID,
		UserName: user.Username,
		CreateAt: user.CreatedAt.Unix(),
	}
}
