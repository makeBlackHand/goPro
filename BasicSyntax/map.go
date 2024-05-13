package main

import "fmt"

func main() {
	var a map[string]string         //map声明不分配空间
	a = make(map[string]string, 10) //使用map前必须make分配空间才可使用
	a["no1"] = "松江"
	a["no2"] = "奥特曼"
	a["no1"] = "打怪兽" //map的key不能重复，value可以重复。重新赋值会覆盖上个值
	a["no3"] = "奥特曼"
	fmt.Println(a)
}
