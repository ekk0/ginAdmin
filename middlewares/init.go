package middlewares

import (
	"github.com/gin-gonic/gin"
)

func InitMiddlewares(c *gin.Context){
	//fmt.Println("中间件...")
	c.Set("username","用户名的中间件") //中间件传值
}