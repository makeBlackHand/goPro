package main

import (
	"errors"
	"fmt"
)

func IsRunYear(a int) (times bool) {
	times = (a%4 == 0 && a%100 != 0) || a%400 == 0
	return
}

func test01() (err error) {
	var year, month, day int
	fmt.Println("请输入年:")
	fmt.Scan(&year)
	fmt.Println("请输入月:")
	fmt.Scan(&month)
	if month <= 0 || month > 12 {
		println("22222")
		err = errors.New("输入异常")
		return
	}
	fmt.Println("请输入日:")
	fmt.Scan(&day)
	flag := IsRunYear(year)
	switch month {
	case 2:
		{
			if flag {
				if day > 0 && day <= 29 {
					println("true")
				} else {
					println("false")
				}
			} else {
				if day > 0 && day <= 28 {
					println("true")
				} else {
					println("false")
				}
			}
		}
	case 4, 6, 9, 11:
		println(30)
	case 1, 3, 5, 7, 10, 12:
		println(31)

	}

	return nil
}

func main() {
	err := test01()
	if err != nil {
		fmt.Println(err)
	}
}
