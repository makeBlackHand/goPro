package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	username, _ := c.Get("username")
	fmt.Println(username)
	v, _ := username.(string)
	con.Success(c)
	c.String(200, v+"首页")
}

func (con UserController) UserNews(c *gin.Context) {
	c.String(http.StatusOK, "用户新闻")
}
