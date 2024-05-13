package main

import (
	"fmt"
	"strconv"
	"time"
)

func testt() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	go testt()

	for i := 1; i <= 10; i++ {
		fmt.Println("main hello csgo" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
