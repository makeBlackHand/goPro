package api

import "github.com/gin-gonic/gin"

type Apicontroller struct {
}

func (con Apicontroller) Index(c *gin.Context) {
	c.String(200, "api首页")
}

func (con Apicontroller) News(c *gin.Context) {
	c.String(200, "api新闻")
}
