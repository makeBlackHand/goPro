package routers

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/3_router_type/controllers/itying"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.ItyingController{}.Index)
		defaultRouters.GET("/news", itying.ItyingController{}.News)
	}
}
