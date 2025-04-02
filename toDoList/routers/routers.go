package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"toDoList/api"
	"toDoList/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()                                        //创建gin的一个实例
	store := cookie.NewStore([]byte("something-very-secret")) //再写入cookie
	r.Use(sessions.Sessions("mysession", store))              //session进行存储

	//写路由
	v1 := r.Group("api/v1") //基础子路由
	{
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		//进行神武验证（中间器）
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("/tasks", api.TasksCreate)      //创建备忘录
			authed.GET("/tasks/:id", api.ShowTask)      //展示一条备忘录
			authed.GET("/tasks", api.ShowTasks)         //展示用户所有备忘录
			authed.PUT("/tasks/:id", api.UpdateTask)    //更新用户备忘录
			authed.POST("/search", api.SearchTask)      // 模糊查询
			authed.DELETE("/tasks/:id", api.DeleteTask) // 删除备忘录
		}
		//这样子写的话，在执行authed后面的路由的时候就会去验证中间件，看你有没有这个权限
	}
	return r
}
