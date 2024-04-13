package main

import "fmt"

func minus(sum *int) {
	*sum += 10
	fmt.Println("sum=", *sum)
	fmt.Println("sum=", sum)
}

func main() {
	var sum int = 10
	fmt.Println("sum=", &sum)
	minus(&sum)
	fmt.Println("sum=", sum)
}
