package routers

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/6_model/controllers/api"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.Apicontroller{}.Index)
		apiRouters.GET("/news", api.Apicontroller{}.News)
	}
}
