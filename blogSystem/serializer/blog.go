package serializer

import "blogSystem/models"

type BlogService struct {
	BlogName     string `json:"blog_name" form:"blog_name" example:"111"`
	BlogContent  string `json:"blog_content" form:"blog_content" example:"222"`
	UserID       uint   `json:"user_id" form:"user_id" example:"1"`
	ID           uint   `json:"id" form:"id" example:"1"`
	CommentCount int    `json:"comment_count"`
}

func BuildBlog(blog models.Blog) BlogService {
	return BlogService{
		BlogName:     blog.BlogName,
		BlogContent:  blog.BlogContent,
		UserID:       blog.UserID,
		ID:           blog.ID,
		CommentCount: blog.CommentCount,
	}
}

func BuildBlogs(blogs []models.Blog) []BlogService {
	result := make([]BlogService, len(blogs))
	for i, blog := range blogs {
		result[i] = BuildBlog(blog)
	}
	return result
}
