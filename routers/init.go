package routers

import (
	"ginAdmin/routers/admin"
	"ginAdmin/routers/api"
	"github.com/gin-gonic/gin"
)
//路由初始化 ,router传参是 gin*Engine 结构体
func InitRouter(router *gin.Engine)  {
	 admin.RoutersApi(router) //引入admin模块
	 api.RoutersApi(router) //引入api模块

}



