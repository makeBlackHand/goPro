package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var p Point = Point{1, 2}
	a = p
	//fmt.Println(a)
	var p1 Point
	p1 = a.(Point) //把空接口转换成对应的类型
	fmt.Println(p1)

	var f float32 = 9.9
	var b interface{}
	b = f
	fmt.Println(b)
	e, flag := b.(float32)
	if flag == true {
		println("转换成功")
		fmt.Println(e)
	} else {
		fmt.Println("转换失败")
	}

}
