package itying

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goPro/gin/6_model/models"
)

type ItyingController struct {
}

func (con ItyingController) Index(c *gin.Context) {
	c.HTML(200, "itying/index.html", gin.H{
		"msg": "ggggggg",
		"t":   1719240271,
	})
}
func (con ItyingController) News(c *gin.Context) {
	fmt.Println(models.UnixToTime(1719240271))
	c.String(200, "itying新闻")
}
