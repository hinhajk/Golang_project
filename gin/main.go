package main

import (
	"gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.GET("/ping", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run(":8080")
	routers.InitRouter(r)
	r.Run(":8080")

}
