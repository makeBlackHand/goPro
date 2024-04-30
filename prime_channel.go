package main

import (
	"fmt"
	"time"
)

func write1(n int, intChan chan int) {
	fmt.Printf("goRoutine %d\n", n)
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func read1(n int, intChan chan int, primeChan chan int, endChan chan bool) {
	fmt.Println("this is goroutine", n)
	//data := <-intChan
	for i := range intChan {
		if isPrime(i) {
			primeChan <- i
		}
	}

	endChan <- true
}

func main() {
	var intChan chan int = make(chan int, 8000)
	var primeChan chan int = make(chan int, 8000)
	var endChan chan bool = make(chan bool, 4)
	start := time.Now()
	write1(0, intChan)

	for i := 0; i < 4; i++ {
		go read1(i, intChan, primeChan, endChan)
	}

	for i := 0; i < 4; i++ {
		<-endChan
	}
	end := time.Now()
	fmt.Println("耗时=", end.Sub(start))
	close(endChan)
	close(primeChan)

	for i := range primeChan {
		println(i)
	}
}
