package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	con.Success(c)
}

func (con UserController) UserNews(c *gin.Context) {
	c.String(http.StatusOK, "用户新闻")
}
