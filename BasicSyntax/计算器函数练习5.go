package main

import "fmt"

func cal(num, num1, num2 int) {
	switch num {
	case 1:
		fmt.Printf("%v=%v+%v", num1+num2, num1, num2)
	case 2:
		fmt.Printf("%v=%v-%v", num1-num2, num1, num2)
	case 3:
		fmt.Printf("%v=%v*%v", num1*num2, num1, num2)
	case 4:
		fmt.Printf("%v=%v/%v", num1/num2, num1, num2)
	}
}
func conf() {
	var num, num1, num2 int
	println("-------------计算器-------------")
	println("1,加法")
	println("2,减法")
	println("3,乘法")
	println("4,除法")
	println("0,退出")
	println("请选择：")
	fmt.Scan(&num)
	println("请输入两个要计算的数")
	fmt.Scan(&num1, &num2)
	cal(num, num1, num2)
}

func main() {
	conf()
}
