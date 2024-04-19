package modle

import (
	"bufio"
	"fmt"
	"os"
)

const path = "clientSystem/modle/log.txt"

type File struct {
}

func (f *File) Store(str string) {
	var err error
	file, err := os.OpenFile(path, os.O_RDWR, 077)
	if err != nil {
		fmt.Println("err = ", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("文件关闭失败")
		}
	}()
	writer1 := bufio.NewWriter(file)
	if nums, err := writer1.WriteString(str); err != nil {
		fmt.Println("写入失败")
	} else {
		fmt.Println("成功， bytes = ", nums)
	}
	func() {
		if err = writer1.Flush(); err != nil {
			fmt.Println("刷新文件失败")
		} else {
			fmt.Println("Success")
		}
	}()
}
