package serializer

import "blogSystem/models"

type CommentService struct {
	UserID      uint   `json:"user_id" form:"user_id" example:"1"`
	ID          uint   `json:"id" form:"id" example:"1"`
	BlogContent string `json:"blog_content" form:"blog_content" example:"222"`
	BlogID      uint   `json:"blog_id" form:"blog_id" example:"1"`
}

func BuildComment(comment models.BlogComment) CommentService {
	return CommentService{
		BlogContent: comment.Content,
		BlogID:      comment.BlogID,
		UserID:      comment.UserID,
		ID:          comment.ID,
	}
}

func BuildComments(comments []models.BlogComment) []CommentService {
	result := make([]CommentService, len(comments))
	for index, comment := range comments {
		result[index] = BuildComment(comment)
	}
	return result
}
