package api

import (
	"blogSystem/service"
	"blogSystem/utils"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateComment(c *gin.Context) {
	var createCommentService service.CommentService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCommentService); err == nil {
		res := createCommentService.CreateCommentService(claim.Id, c.Param("bid"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func GetComments(c *gin.Context) {
	var getCommentsService service.GetCommentsService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getCommentsService); err == nil {
		res := getCommentsService.GetCommentsService(claim.Id, c.Param("bid"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func UserComments(c *gin.Context) {
	var userComments service.GetCommentsService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userComments); err == nil {
		res := userComments.UserCommentsService(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func UpdateComment(c *gin.Context) {
	var updateComment service.CommentService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateComment); err == nil {
		res := updateComment.UpdateCommentService(claim.Id, c.Param("bid"), c.Param("comment_id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func DeleteComment(c *gin.Context) {
	var deleteComment service.CommentService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteComment); err == nil {
		res := deleteComment.BlogDeleteCommentService(claim.Id, c.Param("bid"), c.Param("comment_id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func DeleteComments(c *gin.Context) {
	var deleteComments service.CommentService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteComments); err == nil {
		res := deleteComments.BlogDeleteCommentsService(claim.Id, c.Param("bid"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func UserDeleteComments(c *gin.Context) {
	var deleteComments service.CommentService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteComments); err == nil {
		res := deleteComments.UserDeleteCommentsService(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}
