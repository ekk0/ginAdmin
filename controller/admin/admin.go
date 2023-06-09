package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin struct {

}
func (a Admin) Index(c *gin.Context) {
	 //c.String(http.StatusOK, "***")
	 c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
}



