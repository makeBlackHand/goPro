package main

import "fmt"

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	res := addUpper(10)
	if res != 44 {
		fmt.Println("输出错误", res)
	} else {
		fmt.Println("输出正确", res)
	}
}
