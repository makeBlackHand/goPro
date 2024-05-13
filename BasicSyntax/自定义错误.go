package main

import (
	"errors"
	"fmt"
)

func readConf(name string) (err error) { //返回一个err类型值
	if name == "Config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误") //自己新建异常
	}
}

func test111() {
	err := readConf("Config.inl")
	if err != nil {
		panic(err) //抛出异常，并终止程序
	}
	fmt.Println("test执行")
}

func main() {
	test111()
	fmt.Println("main执行")
}
