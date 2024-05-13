package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "E:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 666)
	//文件类型在Windows下没有作用，只有在Linux有用
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	defer file.Close() //打开文件要记得关闭

	str := "hi,world\n"             //\n表示换行
	writer := bufio.NewWriter(file) //使用缓存写入文件
	for i := 0; i < 10; i++ {
		writer.WriteString(str) //现在只是写在缓存要Flush才能写入文件
	}
	writer.Flush()

}
