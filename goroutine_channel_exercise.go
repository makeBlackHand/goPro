package main

import (
	"fmt"
)

func writes(numChan chan int) {
	for i := 1; i <= 10; i++ {
		//time.Sleep(time.Second)
		//fmt.Println("传数据")
		numChan <- i
	}
	close(numChan)
}

func reads(s int, numChan chan int, resChan chan int, exitChan chan bool) {
	fmt.Println("goroutine ", s)
	var sum int
	numChanLen := cap(numChan)
	for a := 0; a < numChanLen; a++ {
		n, ok := <-numChan
		fmt.Println("n", n, ok)
		if ok {
			sum = 0
			for i := 1; i <= n; i++ {
				sum += i
			}
			//time.Sleep(time.Second)
			//fmt.Println("写数据-")
			resChan <- sum
		}
	}

	exitChan <- true
}

func main() {
	numChan := make(chan int, 10)
	resChan := make(chan int, 10)
	flagChan := make(chan bool, 1)
	writes(numChan)

	for i := 0; i < 1; i++ {
		//time.Sleep(time.Second)
		//fmt.Println("read", i)
		go reads(i, numChan, resChan, flagChan)
	}

	for i := 0; i < 1; i++ {
		<-flagChan
	}
	close(flagChan)
	close(resChan)

	for i := range resChan {
		fmt.Println(i)
	}
}
