package main

import "fmt"

func add1(a int, args ...int) (sum int) {
	sum = a
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func main() {
	sum := add1(1, 2, 3, 3, 4, 5)
	fmt.Println(sum)
}
