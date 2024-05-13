package main

import (
	"fmt"
)

func half(arr *[6]int, a int) {
	var front, end, in int
	front = 0
	end = len(*arr) - 1

	for front <= end {
		in = (front + end) / 2
		if arr[in] == a {
			println("找到了", arr[in], "下标为：", in)
			break
		}
		if a < (*arr)[in] {
			end = in - 1
		} else if a > (*arr)[in] {
			front = in + 1
		}
	}
	println("没找到")
	return
}

func main() {
	var arr = [6]int{1, 4, 6, 12, 46, 79}
	var a int
	fmt.Println("请输入查找的数字：")
	fmt.Scan(&a)
	half(&arr, a)
}
