package service

import (
	"blogSystem/models"
	"blogSystem/serializer"
	"blogSystem/utils"
	"gorm.io/gorm"
)

// UserService 用户注册、登录、登出接口
type UserService struct {
	UserName     string `form:"user_name"`
	PassWord     string `form:"password"`
	Age          int    `form:"age"`
	Sex          string `form:"sex"`
	BlogCount    int    `form:"blog_count"`
	CommentCount int    `form:"comment_count"`
}

// Register 用户注册业务逻辑
func (u *UserService) Register() serializer.Response {
	//	首先查询数据库中是否有该用户
	var user models.User
	var cnt int64
	models.DB.Model(&models.User{}).Where("username = ? ", u.UserName).First(&user).Count(&cnt)
	if cnt > 0 {
		return serializer.Response{
			Status:  403,
			Message: "用户已存在",
		}
	}

	//密码设置是否正确
	user.Username = u.UserName
	user.Sex = u.Sex
	user.Age = u.Age
	err := user.EncryptPassword(u.PassWord)
	if err != nil {
		return serializer.Response{
			Status:  400,
			Message: "密码错误",
		}
	}

	//创建用户
	err1 := models.DB.Create(&user).Error
	if err1 != nil {
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

// Login 用户登录业务逻辑
func (u *UserService) Login() serializer.Response {
	var user models.User
	err := models.DB.Where("username = ? ", u.UserName).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status:  404,
				Message: "用户名不存在,请先注册",
			}
		}

		//如果是其他错误
		return serializer.Response{
			Status:  500,
			Message: "数据库错误",
		}
	}

	//用户存在，验证密码
	flag := user.CheckPassword(u.PassWord)
	if flag == false {
		return serializer.Response{
			Status:  403,
			Message: "密码错误",
		}
	}

	//登陆后，要给用户前发一个token
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

//// Logout 用户退出登录业务逻辑
//func (u *UserService) Logout() serializer.Response {
//
//}

// UserInfo 用户删除、修改、查询接口
type UserInfo struct {
	UserName string `form:"user_name"`
	Age      int    `form:"age"`
	Sex      string `form:"sex"`
}

// ShowInfo 查询用户信息
func (userinfo *UserInfo) ShowInfo(uid string, id uint) serializer.Response {
	var user models.User
	//var cnt int64 = 0
	err := models.DB.First(&user, uid).Error
	if err == nil {
		return serializer.Response{
			Status:  200,
			Message: "用户信息如下",
			Data:    serializer.BuildUser(user),
		}
	} else {
		return serializer.Response{
			Status:  404,
			Message: "用户不存在",
		}
	}
}

// UpdateInfo 更新用户信息
func (userupdate *UserInfo) UpdateInfo(uid string, id uint) serializer.Response {
	var user models.User
	err := models.DB.First(&user, uid).Error
	if err != nil {
		return serializer.Response{
			Status:  404,
			Message: "用户不存在",
		}
	} else {
		user.Username = userupdate.UserName
		user.Age = userupdate.Age
		user.Sex = userupdate.Sex
		models.DB.Save(&user)
		return serializer.Response{
			Status:  200,
			Data:    serializer.BuildUser(user),
			Message: "更新完成",
		}
	}
}
