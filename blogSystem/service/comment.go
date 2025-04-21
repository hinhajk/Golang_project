package service

import (
	"blogSystem/models"
	"blogSystem/serializer"
)

type CommentService struct {
	Content string `json:"comment_content"`
	//`json:"comment_content"`该操作是由于在定义接口时常常将名字大写，而form表单里面是小写
	User models.User
	Blog models.Blog
}

// CreateCommentService 发布评论
func (service *CommentService) CreateCommentService(uid uint, bid string) serializer.Response {
	var comment models.BlogComment
	var user models.User
	models.DB.Where("id = ?", uid).First(&user)
	var blog models.Blog
	models.DB.Where("id = ?", bid).First(&blog)
	comment.Blog = blog
	comment.User = user
	comment.Content = service.Content
	comment.UserID = user.ID
	comment.BlogID = blog.ID
	err := models.DB.Create(&comment).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "评论发布失败",
		}
	} else {
		user.CommentCount++
		models.DB.Save(&user)
		blog.CommentCount++
		models.DB.Save(&blog)
		return serializer.Response{
			Status:  200,
			Message: "评论发布成功",
			Data:    serializer.BuildComment(comment),
		}
	}

}

// GetCommentsService 显示评论接口
type GetCommentsService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

// GetCommentsService 博客评论查询
func (service *GetCommentsService) GetCommentsService(uid uint, bid string) serializer.Response {
	var comments []models.BlogComment
	var count int64
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	err := models.DB.Where("blog_id = ?", bid).Find(&comments).Count(&count).Error
	if count == 0 {
		return serializer.Response{
			Status:  404,
			Message: "该博客没有评论",
		}
	}
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "评论查询失败",
			Error:   err.Error(),
		}
	} else {
		return serializer.Response{
			Status:  200,
			Message: "评论内容如下",
			Data:    serializer.BuildComments(comments),
			Count:   count,
		}
	}
}

// UserCommentsService 用户评论查询
func (service *GetCommentsService) UserCommentsService(uid uint) serializer.Response {
	var userComments []models.BlogComment
	var count int64
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	err := models.DB.Where("user_id = ?", uid).Find(&userComments).Count(&count).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "查询失败",
			Error:   err.Error(),
		}
	} else {
		return serializer.Response{
			Status:  200,
			Message: "用户发布的所有评论如下",
			Data:    serializer.BuildComments(userComments),
			Count:   count,
		}
	}
}

// UpdateCommentService 修改评论
func (service *CommentService) UpdateCommentService(uid uint, bid string, cid string) serializer.Response {
	var comments models.BlogComment
	var count int64 = 0
	models.DB.Where("user_id = ? ", uid).Where("blog_id = ?", bid).Where("id = ? ", cid).Find(&comments).Count(&count)
	if count == 0 {
		return serializer.Response{
			Status:  404,
			Message: "该评论不存在",
		}
	} else {
		comments.Content = service.Content
		models.DB.Save(&comments)
		return serializer.Response{
			Status:  200,
			Message: "更新成功",
			Data:    serializer.BuildComment(comments),
		}
	}
}

// BlogDeleteCommentService DeleteCommentService 删除博客的单条评论
func (service *CommentService) BlogDeleteCommentService(uid uint, bid string, cid string) serializer.Response {
	var comment models.BlogComment
	models.DB.Where("blog_id = ?", bid).Where("id = ?", cid).First(&comment)
	var blog models.Blog
	models.DB.Where("id = ?", bid).First(&blog)
	var user models.User
	models.DB.Where("id = uid").First(&user)
	err := models.DB.Delete(&comment).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "删除失败",
		}
	} else {
		if user.CommentCount == 0 {
			user.CommentCount = 0
		} else {
			user.CommentCount--
		}
		models.DB.Save(&user)
		if blog.CommentCount == 0 {
			blog.CommentCount = 0
		} else {
			blog.CommentCount--
		}
		models.DB.Save(&blog)
		return serializer.Response{
			Status:  200,
			Message: "删除成功",
			Data:    serializer.BuildComment(comment),
		}
	}
}

// BlogDeleteCommentsService 删除博客所有评论
func (service *CommentService) BlogDeleteCommentsService(uid uint, bid string) serializer.Response {
	var comments []models.BlogComment
	var count int64
	models.DB.Where("blog_id = ?", bid).Find(&comments).Count(&count)
	var blog models.Blog
	models.DB.Where("blog_id = ?", bid).First(&blog)
	var user models.User
	models.DB.Where("id = ? ", uid).First(&user)
	err := models.DB.Delete(&comments).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "删除失败",
		}
	} else {
		for _, comment := range comments {
			var user models.User
			models.DB.Where("id = ? ", comment.UserID).First(&user)
			if user.CommentCount > 0 {
				user.CommentCount--
				models.DB.Save(&user)
			}
		}
		blog.CommentCount = 0
		models.DB.Save(&blog)
		return serializer.Response{
			Status:  200,
			Message: "删除成功",
			Data:    serializer.BuildComments(comments),
		}
	}
}

// UserDeleteCommentsService 删除用户所有评论
func (service *CommentService) UserDeleteCommentsService(uid uint) serializer.Response {
	var comments []models.BlogComment
	var user models.User
	models.DB.Where("id = ?", uid).First(&user)
	//var blog models.Blog
	//models.DB.Where("id = ?", bid).First(&blog)
	models.DB.Where("user_id = ?", uid).Find(&comments)
	err := models.DB.Delete(&comments).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "删除失败",
		}
	} else {
		user.CommentCount = 0
		models.DB.Save(&user)
		for _, comment := range comments {
			var blog models.Blog
			models.DB.Model(&models.Blog{}).Where("id = ?", comment.BlogID)
			if blog.CommentCount > 0 {
				blog.CommentCount--
				models.DB.Save(&blog)
			}
		}
		return serializer.Response{
			Status:  200,
			Message: "用户所有评论删除成功",
		}
	}
}
