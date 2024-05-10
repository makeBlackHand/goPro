package main

import (
	"fmt"
	reflect2 "reflect"
)

func testtt(a interface{}) {
	rva := reflect2.ValueOf(a)
	k := rva.Kind()
	t := rva.Type()
	fmt.Printf("kind: %v type:%v\n", k, t)
	s := rva.Interface()
	v := s.(float64)
	fmt.Println("v=", v)
}

func main() {
	var v float64 = 1.2
	testtt(v)
}
