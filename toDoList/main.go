package main

import (
	"toDoList/config"
	"toDoList/routers"
)

func main() { // http://localhost:3000/swagger/index.html
	// 从配置文件读入配置
	config.Init()
	//r := gin.New()
	r := routers.NewRouter()
	r.Run(":3000")
}
