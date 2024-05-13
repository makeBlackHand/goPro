package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(destin string, src string) (zishu int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("err =", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	destFile, err1 := os.OpenFile(destin, os.O_WRONLY|os.O_CREATE, 123)
	if err1 != nil {
		fmt.Println("err1 =", err1)
	}
	defer destFile.Close()
	wrier := bufio.NewWriter(destFile)
	return io.Copy(wrier, reader)
}

func main() {
	srcFile := "C:/emo.png"
	destFile := "E:/emo.png"
	_, err := CopyFile(destFile, srcFile)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("打印完成")
}
