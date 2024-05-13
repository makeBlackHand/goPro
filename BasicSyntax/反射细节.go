package main

import (
	"fmt"
	reflect2 "reflect"
)

func reflect1(a interface{}) {
	rval := reflect2.ValueOf(a)
	fmt.Printf("rval kind is%v\n", rval.Kind())
	rval.Elem().SetInt(20) //相当于指针取值
}

func main() {
	var num int = 10
	reflect1(&num)
	println("num=", num)
}
