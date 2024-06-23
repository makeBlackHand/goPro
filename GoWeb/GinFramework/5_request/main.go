package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   int
	Name string
	Addr string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求地址是:", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求地址后的查询字符串是:", r.URL.RawQuery) //这是地址后的查询字符串
	fmt.Fprintln(w, "请求头的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息是：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性有：", r.Header.Get("Accept-Encoding"))
	//len := r.ContentLength
	////请求体中内容长度
	//body := make([]byte, len)
	//r.Body.Read(body)
	////请求体内容读到body
	//fmt.Fprintln(w, "请求体的内容是:", string(body))
	////浏览器中显示请求体内容
	//r.ParseForm()
	//要先把上面r.body.read方法注释掉，才能获取到值，不然会获取到空白
	//解析表单,在调用r.form之前必须执行该操作
	//fmt.Fprintln(w, "请求参数有:", r.Form)
	//获取请求参数
	//里面有不同类型，根据不同类型调用不同的类型解析
	//fmt.Fprintln(w, "post请求的Form表单的参数有:", r.PostForm)
	//通过调用Formvalue或PostFormValue直接自动获取参数的值
	fmt.Fprintln(w, "URL中user的请求参数有:", r.FormValue("user"))
	fmt.Fprintln(w, "form表单中username的请求参数有:", r.PostFormValue("username"))
}

// 以json的格式返回输出
func TestJsonRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := &User{
		Id:   22,
		Name: "admin",
		Addr: "shanghai",
	}
	data, _ := json.Marshal(user)
	//将user转为json
	w.Write(data)
	//将json格式发给客户端
}

func TestRedire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("location", "https://www.baidu.com")
	//设置响应头的location属性
	w.WriteHeader(302)
	//设置响应状态码
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/TestJsonRes", TestJsonRes)
	http.HandleFunc("/TestRedire", TestRedire)
	http.ListenAndServe("localhost:8080", nil)
}
