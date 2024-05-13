package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [5]int
	var sum, index int
	var large, in int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {

		a := rand.Intn(100)
		arr[i] = a
		fmt.Printf("正序结果为：%d  ", arr[i])
	}
	println()
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("倒叙结果为：%d  ", arr[i])
		sum += arr[i]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 55 {
			println("有55")
			break
		} else {
			println("没有55")
			break
		}
	}
	fmt.Println("平均数为:", float64(sum)/float64(len(arr)))
	for index = 0; index < len(arr); index++ {
		if arr[index] > large {
			large = arr[index]
			in = index
		}
	}
	fmt.Println("最大数是：", large, "最大数下标是：", in)
}
