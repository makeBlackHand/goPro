package routers

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/3_router_type/controllers/admin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admin.UserController{}.Index)
		adminRouters.GET("/news", admin.UserController{}.UserNews)
	}
}
