package main

import (
	"fmt"
)

func write(channel chan int) {
	for i := 0; i < 50; i++ {
		channel <- i * 2
		fmt.Println("write", i)
	}
	close(channel)
}

func read(channel chan int, exit chan bool) {
	for {
		i, v := <-channel
		if v == false {
			break
		}
		fmt.Printf("read:%d\n ", i)
	}
	exit <- true
	close(exit)
}

func main() {
	chan1 := make(chan int, 50)
	chan2 := make(chan bool, 1)
	go write(chan1)
	go read(chan1, chan2)
	v := <-chan2
	if v == true {
		fmt.Println("read完毕！！！！")
	}
}
