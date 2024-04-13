package main

import "fmt"

func main() {
	var pos int
	var neg int
	var i int
	for {
		println("请输入一个数字:")
		fmt.Scanln(&i)
		if i == 0 {
			break
		}
		if i > 0 {
			pos++
			continue
		}
		neg++
	}
	fmt.Printf("正数个数为:%d,负数个数为:%d", pos, neg)
}
