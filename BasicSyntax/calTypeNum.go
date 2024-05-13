package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharNum struct {
	EnCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	file := "E:/实验.txt"
	files, err := os.Open(file)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer files.Close()
	var count CharNum
	reader := bufio.NewReader(files)
	for {
		str, err := reader.ReadString('\n')
		for _, v := range str { //第一个返回的是下标，第二个是值
			switch { //把switch当if使用
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.EnCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("英文字数：%v,数字字数：%v,空格字数：%v,其他字数:%v\n", count.EnCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
