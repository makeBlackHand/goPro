package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 3)
	fmt.Println(ch)
	ch <- 10
	num := 20
	ch <- num
	fmt.Println(num)
	n := <-ch
	println(n)
	ch <- 99
	m := <-ch
	println(m)
	x := <-ch
	println(x)

}
