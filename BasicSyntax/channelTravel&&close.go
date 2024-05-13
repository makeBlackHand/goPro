package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	//intChan <- 300 //管道关闭后不能传值但可以读数据
	a := <-intChan
	fmt.Println(a)
	var chan2 chan int
	chan2 = make(chan int, 100)
	for i := 1; i <= 100; i++ {
		chan2 <- i * 2
	}

	close(chan2) //要关闭管道不然会出现deadlock死锁
	for v := range chan2 {
		fmt.Println(v)
	}
}
