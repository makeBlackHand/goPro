package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("E:/100.txt") //看是否打开成功
	if err != nil {
		fmt.Println("open file err=", err) //file是一个指针返回地址
	}

	fmt.Printf("file=%v", file)

	err = file.Close() //看是否关闭成功
	if err != nil {
		fmt.Println("err close err=", err)
	}
}
