package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var arr = [26]byte{}
	for i := 0; i < 26; i++ {
		arr[i] = byte('A' + i)
		fmt.Printf("%c\n", arr[i]) //打印A-Z
	}
	var index, value int //找出数组中最大值和下标
	var arr1 = [5]int{1, 2, 5, 3, 7}
	for i := 0; i < 4; i++ {
		for j := 1; j < 5; j++ {
			if arr[i] > arr[j] {
				value = arr1[i]
				index = i
			} else {
				value = arr1[j]
				index = j
			}
		}
	}
	fmt.Printf("最大值是:%d,下标是:%d\n", value, index)

	var sun int
	var arr2 = [...]int{2, 3, 41, 1, 31, 1}
	for _, value = range arr2 {
		sun += value
	}
	sum := float64(sun) / float64(len(arr2))
	fmt.Printf("aver:%.2f\n", sum) //求数组中所有数的和的平均数

	var arrs [5]int
	for i := 0; i < len(arrs); i++ {
		arrs[i] = rand.Intn(100)
	}
	fmt.Println("交换前：", arrs)
	for i := 0; i < len(arrs)/2; i++ {
		arrs[i], arrs[len(arrs)-i-1] = arrs[len(arrs)-i-1], arrs[i]
	}
	fmt.Println("交换后：", arrs)
}
