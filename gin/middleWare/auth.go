package middleWare

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 权限检查
func AuthCheck(c *gin.Context) {
	userId, _ := c.Get("user_id")
	UserName, _ := c.Get("user_name")
	fmt.Println(userId, UserName)
}

var token = "123456"

func TokenCheck(c *gin.Context) {
	accessToken := c.Request.Header.Get("access_toke")
	if accessToken != token {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "middleWare_check error",
		})
	}
	c.AbortWithError(http.StatusInternalServerError, errors.New("middleWare_check error"))
	c.Set("user_name", "nick")
	c.Set("user_id", "10001")
	c.Next()

}
