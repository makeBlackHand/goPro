package itying

import "github.com/gin-gonic/gin"

type ItyingController struct {
}

func (con ItyingController) Index(c *gin.Context) {
	c.HTML(200, "itying/index.html", gin.H{
		"msg": "ggggggg",
	})
}
func (con ItyingController) News(c *gin.Context) {
	c.String(200, "itying新闻")
}
