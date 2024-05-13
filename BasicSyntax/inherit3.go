package main

import "fmt"

type monster struct {
	name string
}
type D struct {
	monster
	int //可以直接定义基本元素类型，只能定义一次
}

func main() {
	var d D
	d.name = "哈哈哈"
	d.int = 99
	fmt.Println(d)
}
