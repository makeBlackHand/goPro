package main

import "fmt"

func alert(map1 map[int]int) {
	map1[1] = 99
}

type Stu struct {
	name    string
	age     int
	address string
}

func main() {
	map1 := make(map[int]int) //map可以自动扩容
	map1[1] = 10
	map1[2] = 100
	map1[0] = 11
	map1[10] = 123
	alert(map1) //map是引用类型，再定义会覆盖
	fmt.Println(map1)

	students := make(map[string]Stu, 10)
	stu1 := Stu{"张三", 100, "北京"}
	stu2 := Stu{"李四", 99, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	for s, stu := range students {
		fmt.Println("学生编号是：", s)
		fmt.Println("学生年龄是：", stu.age)
		fmt.Println("学生地址是：", stu.address)
	}
}
