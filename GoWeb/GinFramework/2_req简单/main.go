package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //方法重写
	fmt.Fprintln(w, "通过详细创建的处理器处理请求")
}

// handlerfun是自动把函数转换成处理器而handle是本来就是处理器直接用!!!
func main() {
	m := MyHandler{}
	server := http.Server{ //详细配置处理器
		Addr:        ":8080",
		Handler:     &m,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
	//http.Handle("/m", &m) //需要自己定义结构体并重写此方法
	//创建路由
	//http.ListenAndServe("localhost:8080", nil)
}
