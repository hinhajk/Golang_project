package api

import (
	"github.com/gin-gonic/gin"
	"toDoList/service"
)

// 创建userResgiter接口（用户注册接口）
// 先声明一个User服务对象
func UserRegister(c *gin.Context) {
	var userRegister service.UserService // 定义一个user服务变量
	// 对他进行一个绑定
	if err := c.ShouldBind(&userRegister); //绑定后一个服务对象，将绑定的值传递过来
	err == nil {
		res := userRegister.Register() //对其执行注册的一个方法
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// 用户登录接口
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
