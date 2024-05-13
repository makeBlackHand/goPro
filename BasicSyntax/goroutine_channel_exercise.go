package main

import "fmt"

func write100(intChan chan int) {
	for i := range 100 {
		intChan <- i
	}
	close(intChan)
	fmt.Println("Finished.")
}

func fib64(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func procecss100(s int, intChan chan int, resChan chan int, exitChan chan bool) {
	fmt.Println("Process ", s)
	numChan := len(intChan)
	for i := 0; i < numChan; i++ {
		resChan <- fib64(<-intChan)
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 100)
	resChan := make(chan int, 100)
	endChan := make(chan bool, 4)

	write100(intChan)

	for i := 0; i < 4; i++ {
		go procecss100(i, intChan, resChan, endChan)
	}

	for i := 0; i < 4; i++ {
		<-endChan
	}
	close(endChan)
	close(resChan)

	for i := 0; i < 100; i++ {
		println(<-resChan)
	}
}
