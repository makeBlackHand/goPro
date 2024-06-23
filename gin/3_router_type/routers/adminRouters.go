package routers

import "github.com/gin-gonic/gin"

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "admin首页")
		})
		adminRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "admin新闻")
		})
	}
}
