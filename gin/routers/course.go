package routers

import (
	"gin/web"
	"github.com/gin-gonic/gin"
)

func initCourse(r *gin.Engine) {
	//course := r.Group("/course") //定义版本
	v1 := r.Group("/v1") // 定义资源
	v1.POST("/course", web.Create)
	v1.GET("/courses", web.Get)
	v1.DELETE("/course", web.DELETE)
	v1.PUT("/course", web.Edit)
}
