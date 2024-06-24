package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func initmiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-我是中间件")
	c.Next()
	//调用该请求（get）的剩余处理程序
	fmt.Println("2-中间件")
	end := time.Now().UnixNano()
	fmt.Println("执行时间为:", end-start)
	//计算一个get的运行时间
}

func initmiddlewareOne(c *gin.Context) {
	fmt.Println("1-我是中间件--1")
	c.Next()
	fmt.Println("2-中间件--1")
}
func initmiddlewareTwo(c *gin.Context) {
	fmt.Println("1-我是中间件--2")
	c.Next()
	fmt.Println("2-中间件--2")
}
func main() {
	r := gin.Default()
	r.Use(initmiddleware)
	//这是全局中间件

	//middleware就是中间件
	r.GET("/", initmiddleware, func(c *gin.Context) { //这些都是路由中间件
		fmt.Println("我是首页")
		c.String(200, "首页")
	})
	r.GET("/news", initmiddlewareOne, initmiddlewareTwo, func(c *gin.Context) {
		fmt.Println("我是新闻")
		c.String(200, "新闻")
	})
	r.GET("/login", func(c *gin.Context) {
		fmt.Println("我是登录")
		c.String(200, "登录")
	})
	r.Run()
}
