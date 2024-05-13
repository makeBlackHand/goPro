package main

import "fmt"

var (
	res2 = func(c, d int) int {
		return c * d
	}
)

func main() {
	res := func(a, b int) int { //匿名函数只能调用一次
		return a + b
	}(10, 20)
	fmt.Println("res=", res)
	res1 := func(c, d int) int {
		return c - d
	}
	sum := res1(20, 10)
	println("sum=", sum)
	sum1 := res2(10, 3)
	println("sum1=", sum1)
}
