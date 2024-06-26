package routers

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/3_router_type/controllers/api"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.Apicontroller{}.Index)
		apiRouters.GET("/news", api.Apicontroller{}.News)
	}
}
