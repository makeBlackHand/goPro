package main

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/3_router_type/routers"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	routers.AdminRoutersInit(r)
	//routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)
	r.Run(":8010")
}
