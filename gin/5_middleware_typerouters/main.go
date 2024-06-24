package main

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/5_middleware_typerouters/routers"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)
	r.Run(":8080")
}
