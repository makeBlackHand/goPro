package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	num1 int
}
type B struct {
	num1 int
}
type Monster struct {
	Name string `json:"name"` //将首字母小写
	Age  int    `json:"age"`  //这就是struct tag
}

func main() {
	var a A
	var b B
	a = A(b) //不同类型结构体不能赋值,强转要里面属性名一样属性一样
	//type 转换名称后它会认为转换后的名称是一种新的类型要强转使用
	fmt.Println(a, b)

	m := Monster{"妲己", 99}
	js, err := json.Marshal(m) //将m变量序列化为json格式字串   这就是反射
	if err != nil {
		fmt.Println("错误")
	} else {
		fmt.Println("js=", string(js)) //返回的是[]byte切片
	}
}
