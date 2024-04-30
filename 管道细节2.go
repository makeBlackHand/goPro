package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	stringChan := make(chan string, 5)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	for i := 0; i < 5; i++ {
		stringChan <- fmt.Sprintf("%d", i)
	}
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从字符串int取到值：%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从字符串stringChan取到值：%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("都取不到了！！")
			return
		}
	}
}
