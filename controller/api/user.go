package api

import (
	"fmt"
	"ginAdmin/logic/api"
	"github.com/gin-gonic/gin"
	"net/http"
)
type UserApi struct {
	BaseApi
}

//继承baseapi
func (b UserApi) Index(c *gin.Context) {
	//逻辑层查询
	  userLine :=logic.GetUserById(1)

	 //username,_ := c.Get("username")
	  fmt.Println("----------------",userLine)
	  c.JSON(http.StatusOK,gin.H{
	  	"result":userLine,
	  })
	// c.String(http.StatusOK, "")
	//b.Ok()

}
func (b UserApi) Add(c *gin.Context)  {

	c.String(http.StatusOK, "index")
}

func (b UserApi) Del(c *gin.Context)  {
	c.String(http.StatusOK, "index")
}

