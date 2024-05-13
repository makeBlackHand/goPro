package main

import (
	"fmt"
)

func main() {
	var arr [5]int = [5]int{10, 15, 24, 46, 120}
	var arr1 [6]int
	var num, i int
	println("请输入一个要输入的数字")
	fmt.Scan(&num)
	for i = 0; i < len(arr); i++ {
		arr1[i] = arr[i]
	}
	arr1[i] = num
	fmt.Println("乱序：", arr1)
	for j := 0; j < len(arr1); j++ {
		for k := 0; k < len(arr1)-1; k++ {
			if arr1[k] > arr1[k+1] {
				arr1[k], arr1[k+1] = arr1[k+1], arr1[k]
			}
		}
	}
	fmt.Println(arr1)
}
