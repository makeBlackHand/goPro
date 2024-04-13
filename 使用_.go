package main

import "fmt"

func sumAndSub(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	a, _ := sumAndSub(10, 5)
	fmt.Printf("a=%d", a)
}
