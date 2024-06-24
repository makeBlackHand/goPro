package routers

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/6_model/controllers/admin"
	"goPro/gin/6_model/middleWare"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middleWare.InitMiddleWare)
	{
		adminRouters.GET("/", admin.UserController{}.Index)
		adminRouters.GET("/news", admin.UserController{}.UserNews)
	}
}
