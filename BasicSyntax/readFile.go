package main

import (
	"fmt"
	"os"
)

func main() {
	str, err := os.ReadFile("E:/100.txt") //直接调用os里的方法
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("file = ", string(str))
}
