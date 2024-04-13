package main

import "fmt"

var sum = test()

func init() {
	fmt.Println("init....")
}

func test() int {
	fmt.Println("test...")
	return 90
}

func main() {
	fmt.Println("main...")
	fmt.Println("sum=", sum)
}
