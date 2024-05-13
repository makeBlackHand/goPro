package main

import (
	"errors"
	"fmt"
)

func err1(name string) (err error) {
	if name != "hello" {
		return errors.New("输入异常")
	}
	return nil
}

func testl() {
	e := err1("1111")
	if e != nil {
		panic(e)
	}
}

func err() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		return
	}()
	a := 10
	b := 0
	c := a / b
	fmt.Println(c)
}

func main() {
	testl()
}
