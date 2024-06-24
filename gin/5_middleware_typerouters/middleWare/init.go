package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleWare(c *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	c.Set("username", "张三")
	//在中间件里面设置数据

	//定义一个goroutine统计日志
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("DONE！ in path" + cCp.Request.URL.Path)
	}()
}
