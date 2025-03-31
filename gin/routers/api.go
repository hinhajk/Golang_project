package routers

import (
	"gin/middleWare"
	"gin/web"
	"github.com/gin-gonic/gin"
)

func initApi(r *gin.Engine) {
	//var i int = 18
	// http://localhost:8080/api
	api := r.Group("/api") //定义版本
	//如果在api这里加一个中间件，则只对api生效

	v1 := api.Group("/v1", middleWare.TokenCheck, middleWare.TokenCheck) //定义资源
	v1.GET("/ping", web.Ping)                                            //定义请求方法，web.Ping是传入方法，不是函数调用
	v1.POST("/login", web.Login)
	v1.POST("/register", web.Register)
	//在上述例子中，访问的URL则是http://localhost:8080/api/v1/ping
	//http://localhost:8080/api/v1/register
	//http://localhost:8080/api/v1/login
}
