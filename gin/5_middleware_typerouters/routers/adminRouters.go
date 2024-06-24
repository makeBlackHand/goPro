package routers

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/5_middleware_typerouters/controllers/admin"
	"goPro/gin/5_middleware_typerouters/middleWare"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middleWare.InitMiddleWare)
	{
		adminRouters.GET("/", admin.UserController{}.Index)
		adminRouters.GET("/news", admin.UserController{}.UserNews)
	}
}
