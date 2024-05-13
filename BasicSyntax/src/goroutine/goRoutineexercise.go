package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex //互斥锁
)

func jieCheng(n int) {
	var sum int = 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	lock.Lock()    //加锁
	myMap[n] = sum //多个协程并发写map
	lock.Unlock()  //加锁后要解锁
}
func main() {
	for i := 1; i <= 200; i++ {
		go jieCheng(i)
	}
	for index, value := range myMap {
		fmt.Printf("map[%v]=%v\n", index, value)
		time.Sleep(time.Second)
	}
}
