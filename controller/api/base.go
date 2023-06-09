package api

import "github.com/gin-gonic/gin"

type BaseApi struct {


}
func (b BaseApi) Ok(c *gin.Context) {


	c.String(200,"okkkk")
}