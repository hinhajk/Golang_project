package main

import (
	"blogSystem/config"
	"blogSystem/routers"
)

func main() {
	config.Init()
	r := routers.NewRouter()
	r.Run(":3000")
}
