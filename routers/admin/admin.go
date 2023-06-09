package admin

import (
	"ginAdmin/controller/admin"
	"github.com/gin-gonic/gin"

)

func RoutersApi(r *gin.Engine) {

	ad := r.Group("/admin")
	{

		r.LoadHTMLGlob("templates/**/**/*")
		ad.GET("/admin/index",admin.Admin{}.Index)

	}
	{

	}



}
