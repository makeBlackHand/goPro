package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World", i)
	}
}

func Test() {
	defer func() { //协程可以捕获异常
		if err := recover(); err != nil {
			fmt.Println("test错误：", err)
		}
	}()
	var myMap map[int]string
	//myMap = make(map[int]string)
	myMap[0] = "张三"
}
func main() {
	go sayHello()
	go Test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("main", i)
	}
}
