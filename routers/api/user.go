package api

import (
	"ginAdmin/controller/api"
	"ginAdmin/middlewares"
	"github.com/gin-gonic/gin"
)


func RoutersApi(r *gin.Engine) {
	//中间件的使用
	ap := r.Group("/api",middlewares.InitMiddlewares)
	{	//用户
		ap.GET("/user/index",api.UserApi{}.Index)
		ap.GET("/user/delete",api.UserApi{}.Del)
		ap.GET("/user/add",api.UserApi{}.Add)
	}
	{ 	//文章
		ap.GET("/article/index",api.Artice{}.Index)

	}
	{  //...

	}


}