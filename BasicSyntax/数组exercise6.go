package main

import "fmt"

func main() {
	var nums string
	var num = [6]string{"AA", "BB", "CC", "DD", "EE", "AA"}
	println("请输入要查的:")
	fmt.Scan(&nums)
	for i := 0; i < len(num); i++ {
		if nums == num[i] {
			fmt.Println("找到了 ", nums, "  下标是：", i)
		}
	}
}
