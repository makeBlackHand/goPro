package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己的多路复用处理器处理请求！", r.URL.Path)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World22222", r.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	//自己创建多路复用器

	//http会注册到默认,换成mux回存到nux
	mux.HandleFunc("/", handler) //+Func是把普通方法装换成处理器（更让方便快捷）
	http.HandleFunc("/", handler2)

	//创建路由
	http.ListenAndServe("localhost:8080", mux) //用自己的多路复用器
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	data, _ := json.Marshal(&Info{
//		Name: "kanna",
//		Age:  24,
//	})
//	w.Write(data)
//}
//
//func main() {
//	http.HandleFunc("/", handler)
//	http.ListenAndServe("localhost:8080", nil)
//
//}
