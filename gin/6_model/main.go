package main

import (
	"github.com/gin-gonic/gin"
	"goPro/gin/6_model/models"
	"goPro/gin/6_model/routers"
	"html/template"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	r.LoadHTMLGlob("templates/**/*")
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)
	r.Run(":8080")
}
