package main

import (
	"github.com/gin-gonic/gin"
	"ginAdmin/routers"
)

func main() {
	// 1.创建路由
	r :=gin.Default()
	routers.InitRouter(r)
	//运行8000
	//引入静态文件
	r.Static("/assets", "./assets")

	r.Run(":8000")
}