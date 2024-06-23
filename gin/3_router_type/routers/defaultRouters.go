package routers

import "github.com/gin-gonic/gin"

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.HTML(200, "default/index.html", gin.H{
				"msg": "ggggggg",
			})
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "default新闻")
		})
	}
}
