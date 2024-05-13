package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("E:/100.txt")
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("file = ", file)

	defer file.Close()

	reader := bufio.NewReader(file) //用到缓冲读文件
	for {
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF {
			break
		}
	}

}
