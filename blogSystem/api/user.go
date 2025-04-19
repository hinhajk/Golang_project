package api

import (
	"blogSystem/service"
	"blogSystem/utils"

	//"blogSystem/utils"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}

//UserLogout 用户退出登录接口
//func UserLogout(c *gin.Context) {
//	var userLogout service.UserService
//	if err := c.ShouldBind(&userLogout); err == nil {
//		res := userLogout.Logout()
//		c.JSON(res, 200)
//	} else {
//		c.JSON(400, gin.H{"error": err.Error()})
//	}
//}

// UserShow 查询用户信息
func UserShow(c *gin.Context) {
	var listTasks service.UserInfo

	//先进行身份验证
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTasks); err == nil {
		res := listTasks.ShowInfo(c.Param("id"), claim.Id) //后面的id为备忘录的id,前面的id为用户id可有cookie获得
		c.JSON(200, res)
	} else {
		//有错误的话，返回并打印日志
		logging.Error(err)
		c.JSON(400, err)
	}
}

// UserUpdate 修改用户信息
func UserUpdate(c *gin.Context) {
	var updateUser service.UserInfo

	//先进行身份验证
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateUser); err == nil {
		res := updateUser.UpdateInfo(c.Param("id"), claim.Id) //后面的id为备忘录的id,前面的id为用户id可有cookie获得
		c.JSON(200, res)
	} else {
		//有错误的话，返回并打印日志
		logging.Error(err)
		c.JSON(400, err)
	}
}
