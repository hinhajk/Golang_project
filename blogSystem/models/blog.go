package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	BlogName     string `gorm:"type:varchar(255);not null" json:"blog_name"`
	BlogContent  string `gorm:"type:longtext;not null" json:"blog_content"`
	User         User   `gorm:"ForeignKey:UserID; references:id" json:"user"`
	UserID       uint   `gorm:"not null" json:"user_id"`
	CommentCount int    `gorm:"default:0" json:"comment_count"`
}
