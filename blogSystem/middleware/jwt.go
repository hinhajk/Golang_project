package middleware

import (
	"blogSystem/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// JWT 中间件验证token
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		//var data interface{}
		//鉴权一般都放在Authoritarian这里
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404 // token为空
		} else {
			claim, err := utils.ParseToken(token)

			//对错误进行处理
			if err != nil {
				code = 403 //无权限，token是假的
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //token已过期
			}
		} //token为空则返回400，不为空则对token进行解析

		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
