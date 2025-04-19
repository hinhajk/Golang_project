package routers

// 该文件是路由定义文件
import (
	"blogSystem/api"
	"blogSystem/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))   //写入cookie
	r.Use(sessions.Sessions("mysession", store)) //用session存储

	//定义url
	v1 := r.Group("/api/v1")
	{
		//用户操作的URL
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)

		authed := v1.Group("/")
		////对博客进行操作时首先要对用户进行身份验证，用中间件来完成
		authed.Use(middleware.JWT())
		{
			authed.GET("/user/show/:id", api.UserShow)
			authed.PUT("/user/update/:id", api.UserUpdate)

			authed.POST("/blogs/create", api.CreateBlog)       //增加博客
			authed.GET("/blogs/show/:id", api.GetBlog)         //查询单挑博客
			authed.GET("/blogs/shows/:uid", api.GetBlogs)      //查询用户发布的所有博客
			authed.PUT("/blogs/update/:id", api.UpdateBlog)    //更新博客内容
			authed.POST("/blogs/search", api.SearchBlogs)      //模糊查询博客
			authed.DELETE("/blogs/delete/:id", api.DeleteBlog) //删除单个博客
			authed.DELETE("/blogs/delete", api.DeleteBlogs)    //删除用户发布的所有博客内容

			authed.POST("/blogs/:bid/comment/create", api.CreateComment)                //增加评论
			authed.GET("/blogs/:bid/comments", api.GetComments)                         //查询该条博客的所有评论
			authed.GET("/blogs/comments", api.UserComments)                             //查询用户发布的评论
			authed.PUT("/blogs/:bid/comments/update/:comment_id", api.UpdateComment)    //更新该条评论内容
			authed.DELETE("/blogs/:bid/comments/delete/:comment_id", api.DeleteComment) //删除单条评论
			authed.DELETE("/blogs/:bid/comments/delete", api.DeleteComments)            //删除博客的所有评论
			authed.DELETE("/blogs/user/comments/delete", api.UserDeleteComments)        //删除用户发布的所有评论
		}
	}
	return r
}
