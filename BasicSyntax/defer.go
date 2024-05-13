package main

import (
	"fmt"
)

func defer1(a, b int) (sum int) {
	defer fmt.Println(a)
	defer fmt.Println(b)
	a++
	b++
	sum = a + b
	fmt.Println(sum)
	return
}
func main() {
	count := defer1(10, 11)
	fmt.Println(count)
}
