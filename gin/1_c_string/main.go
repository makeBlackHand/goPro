package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	//配置模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "首页")
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})
	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好gin--2",
		})
	})
	r.GET("/json3", func(c *gin.Context) {
		a := &Article{
			Title:   "fff",
			Desc:    "fadf",
			Content: "gfad",
		}
		c.JSON(200, a)
	})

	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:   "fff",
			Desc:    "fadf",
			Content: "gfad",
		}
		c.JSONP(200, a)
		//jsonp?callback=xxx
		//xxx({"title":"fff","desc":"fadf","content":"gfad"});

	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, map[string]interface{}{ //或者用gin.H
			"success":  true,
			"messsage": "我是xml",
		})
	})
	r.GET("/news", func(c *gin.Context) {
		//注意//配置模板文件/要与目录一致
		//	r.LoadHTMLGlob("templates/*")
		c.HTML(http.StatusOK, "news.html", gin.H{ //或者用gin.H
			"title": "我是后台数据", //把后台数据传到前端
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		//注意//配置模板文件/要与目录一致
		//	r.LoadHTMLGlob("templates/*")
		c.HTML(http.StatusOK, "goods.html", gin.H{ //或者用gin.H
			"title": "我是商品页面", //把后台数据传到前端
			"price": 20,
		})
	})
	r.Run(":8080")
}
