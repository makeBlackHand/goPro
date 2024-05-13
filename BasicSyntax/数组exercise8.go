package main

import "fmt"

func arrTest01(arr [5]int) {
	var min1, max1, minIndex, maxIndex int
	min1 = arr[0]
	minIndex = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max1 {
			max1 = arr[i]
			maxIndex = i
		}
		if arr[i] < min1 {
			min1 = arr[i]
			minIndex = i
		}

	}
	fmt.Printf("最大值为：%d,下标为：%d,最小值为：%d,下标为:%d", max1, maxIndex, min1, minIndex)
}

func main() {
	var arr = [5]int{12, 34, 123, 470, 246}
	arrTest01(arr)
}
