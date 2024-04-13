package main

import "fmt"

type Person3 struct {
	Name string
	Age  int
}

func main() {
	var p = Person3{"ad", 12} //第一种赋值方式（顺序赋值方式）

	var p1 = Person3{ //第二种赋值方式（比第一种可以不按顺序赋值）
		Age:  12,
		Name: "fad",
	}

	var p2 = &Person3{"fas", 345} //第三种赋值方式，p2通过地址找到结构体
	//这时候p2是个地址
	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(*p2)
}
