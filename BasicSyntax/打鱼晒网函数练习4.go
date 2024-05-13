package main

import (
	"fmt"
)

func IsRunYear1(a int) (times bool) {
	times = (a%4 == 0 && a%100 != 0) || a%400 == 0
	return
}

func test4(year, month, day int) (totalDay int) {
	dayy := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var i int
	for i = 1990; i < year; i++ {
		if IsRunYear1(i) {
			totalDay += 366
		} else {
			totalDay += 365
		}
	}
	for j := 1; j < month; j++ {
		if IsRunYear1(i) {
			dayy[2] = 29
		}
		totalDay += dayy[j]
	}
	totalDay += day
	return
}

func main() {
	var year, month, day int
	println("请输入年月日：")
	fmt.Scan(&year, &month, &day)
	test4(year, month, day)
	num := test4(year, month, day)
	if num%5 > 0 && num%5 <= 3 {
		println("在打鱼")
	} else {
		println("在晒网")
	}
}
