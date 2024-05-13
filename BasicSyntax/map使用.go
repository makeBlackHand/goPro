package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string)
	fmt.Println(a)

	b := make(map[string]string)
	b["no1"] = "上海"
	b["no2"] = "北京"
	fmt.Println(b)

	c := map[string]string{
		"no1": "hello",
		"no2": "world",
	}
	fmt.Println(c)

	studentmap := make(map[string]map[string]string) //学号可以保存多个值
	studentmap["no1"] = make(map[string]string, 2)
	studentmap["no1"]["sex"] = "男"
	studentmap["no1"]["name"] = "张三"
	studentmap["no2"]["sex"] = "女"
	studentmap["no2"]["name"] = "孙尚香"
	fmt.Println(studentmap)

}
