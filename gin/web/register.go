package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerReq struct {
	Username string `form:"username" binding:"required"`
	Pwd      string `form:"password" binding:"required"`
	Phone    string `form:"phone" binding:"required, e164"`   //必填，且格式为e164的格式
	Email    string `form:"email" binding:"omitempty, email"` //前面为空的情况下必填
}

func Register(c *gin.Context) {
	req := &registerReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}
