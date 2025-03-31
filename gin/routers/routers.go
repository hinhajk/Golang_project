package routers

import "github.com/gin-gonic/gin"

// 1.初始化路由
func InitRouter(r *gin.Engine) {
	// 初始化课程路由
	// 初始化API 路由（不需要鉴权）
	r.Use() //在此处的话，则意味着全局生效
	initApi(r)
	initCourse(r)
}
