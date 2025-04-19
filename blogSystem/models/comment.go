package models

import "gorm.io/gorm"

type BlogComment struct {
	gorm.Model
	Content string `gorm:"type:longtext" json:"comment_content"`
	User    User   `gorm:"ForeignKey:UserID; references:id" json:"user"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	Blog    Blog   `gorm:"ForeignKey:BlogID; references:id" json:"blog"`
	BlogID  uint   `gorm:"not null" json:"blog_id"`
}
