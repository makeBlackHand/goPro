package main

import "fmt"

type Cat struct {
	Name  string
	Color string
	Age   int
	ptr   *int  //没有分配空间都是nil
	slice []int //指针，切片和map要先分配空间才能使用
	map1  map[string]string
}

func main() {
	var cat1 Cat
	cat1.Name = "花泽类"
	cat1.Age = 99
	cat1.Color = "花不溜秋"
	fmt.Println("第一只猫：", cat1)

	cat1.slice = make([]int, 3)
	cat1.slice[0] = 99

	cat1.map1 = make(map[string]string)
	cat1.map1["no1"] = "张三"
	fmt.Println(cat1)
}
