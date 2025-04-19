package serializer

import "blogSystem/models"

type User struct {
	ID           uint   `json:"id" form:"id" example:"1"`
	UserName     string `json:"user_name" form:"user_name" example:"FanOne"`
	Sex          string `json:"sex" form:"sex" example:"male"`
	Age          int    `json:"age" form:"age" example:"18"`
	BlogCount    int    `json:"blog_count" form:"blog_count" example:"0"`
	CommentCount int    `json:"comment_count" form:"comment_count" example:"0"`
}

// BuildUser 序列化用户
func BuildUser(user models.User) User {
	return User{
		ID:           user.ID,
		UserName:     user.Username,
		Sex:          user.Sex,
		Age:          user.Age,
		BlogCount:    user.BlogCount,
		CommentCount: user.CommentCount,
	}
}
