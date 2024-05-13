package main

import "fmt"

func input(channel chan int) {
	for i := 1; i < 1001; i++ {
		channel <- i
	}
	close(channel)
}

func output(channel chan int, resChan chan int, flag chan bool) {
	long := 1000
	for i := 1; i <= long; i++ {
		n := <-channel
		sum := 0
		for i := 1; i <= n; i++ {
			sum += i
		}
		resChan <- sum
	}
	flag <- true
}

func main() {
	ch := make(chan int, 1000)
	reChan := make(chan int, 1000)
	flahChan := make(chan bool, 8)

	input(ch)

	for i := 0; i < 8; i++ {
		go output(ch, reChan, flahChan)
	}

	for i := 0; i < 8; i++ {
		<-flahChan
	}

	close(flahChan)
	close(reChan)

	for i2 := range reChan {
		fmt.Println(i2)
	}
}
