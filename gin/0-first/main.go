package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//创建一个默认的路由引擎
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值%v", "你好gin")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "新闻页")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "主要用于增加数据")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "主要用于编辑数据")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "主要用于删除数据")
	})
	r.Run(":8010") //启动一个web服务
}
