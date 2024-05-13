package main

import "fmt"

func main() {
	//声明为只写
	var chan1 chan<- int
	chan1 = make(chan int, 1)
	chan1 <- 20
	fmt.Println()

	//声明为只读
	var chan2 <-chan int
	num2 := <-chan2
	fmt.Println(num2)
}
