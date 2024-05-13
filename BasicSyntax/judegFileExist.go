package main

import (
	"fmt"
	"os"
)

func pathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //证明文件或目录已经存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func main() {
	file1Path := "E:/abc.txt"
	file2Path := "E:/111.txt"
	b, e := pathExist("E:/hhhh.txt")
	fmt.Printf("bool=%v,error=%v\n", b, e)
	data, err := os.ReadFile(file1Path)
	if err != nil {
		fmt.Println("读取错误", err)
		return
	}
	fmt.Println(string(data))
	err1 := os.WriteFile(file2Path, data, 123)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
	}
}
