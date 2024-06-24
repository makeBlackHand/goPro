package routers

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/5_middleware_typerouters/controllers/itying"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.ItyingController{}.Index)
		defaultRouters.GET("/news", itying.ItyingController{}.News)
	}
}
