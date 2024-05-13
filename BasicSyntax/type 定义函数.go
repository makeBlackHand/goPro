package main

import "fmt"

type myFUun func(int, int) int

func myFun2(myFun1 myFUun, a int, n int) int {
	return myFun1(a, n)
}

func add(a, b int) int {
	return a + b
}

func main() {
	type myInt int
	var num1 myInt
	var num2 int
	num1 = 10
	num2 = 20
	num2 = int(num1)
	fmt.Printf("num1=%d,num2=%d\n", num1, num2)
	fmt.Println(myFun2(add, 1, 2))
}
