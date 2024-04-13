package main

import "fmt"

func fib(a int) []uint64 {
	f := make([]uint64, a)
	f[0] = 1
	f[1] = 1
	for i := 2; i < a; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

func main() {
	var a int
	println("请输入数字:")
	fmt.Scan(&a)
	sun := fib(a)
	fmt.Println(sun)
}
