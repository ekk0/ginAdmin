package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Artice struct {

}

func (a Artice) Index(c *gin.Context) {
	c.String(http.StatusOK, "user")

}

