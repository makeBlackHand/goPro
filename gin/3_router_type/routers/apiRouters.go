package routers

import "github.com/gin-gonic/gin"

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "api首页")
		})
		apiRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "api新闻")
		})
	}
}
