package service

import (
	"gorm.io/gorm"
	"toDoList/pkg/utils"

	//"github.com/gin-gonic/gin"
	"toDoList/models"
	"toDoList/serializer"
)

type UserService struct {
	UserName string `form:"user_name"`
	PassWord string `form:"password"`
}

// 定义一个Userservice方法，定义其注册业务逻辑
func (u *UserService) Register() serializer.Response {
	var user models.User
	//首先验证一下数据库里面有没有这个人
	var cnt int64
	models.DB.Model(&models.User{}).Where("username = ?", u.UserName).First(&user).Count(&cnt)
	if cnt == 1 { // 用户已经存在
		return serializer.Response{
			Status:  400,
			Message: "用户已存在",
		}
	}
	user.Username = u.UserName
	err := user.SetPassWord(u.PassWord)
	if err != nil {
		return serializer.Response{
			Status:  400,
			Message: "密码错误",
		}
	}
	//创建用户
	if err := models.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status:  500,
			Message: "数据库操作错误",
		}
	}
	return serializer.Response{
		Status:  200,
		Message: "用户注册成功",
	}
}

func (u *UserService) Login() serializer.Response {
	var user models.User
	//首先验证用户名是否存在，不存在则用户名错误
	//var cnt int64 = 0
	err := models.DB.Where("username = ?", u.UserName).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status:  400,
				Message: "用户名不存在,请先注册",
			}
		}

		//如果是其他错误
		return serializer.Response{
			Status:  500,
			Message: "数据库错误",
		}
	}

	//用户名存在，验证密码
	flag := user.CheckPassWord(u.PassWord)
	if !flag {
		return serializer.Response{
			Status:  400,
			Message: "密码错误",
		}
	}

	//密码验证通过

	//发一个token，为了其他功能身份验证所给前端存储的，例如：创建一个备忘录这个功能就需要token，不然不知道是谁创建的
	token, err1 := utils.GenerateToken(user.ID, u.UserName, u.PassWord)
	if err1 != nil {
		return serializer.Response{
			Status:  500,
			Message: "Token签发错误",
		}
	}
	return serializer.Response{
		Status:  200,
		Data:    serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Message: "登陆成功",
	}
}
