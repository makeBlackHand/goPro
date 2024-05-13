package main

import (
	"fmt"
	reflect2 "reflect"
)

type M struct {
	Num1 int
	Num2 int
}

func (m M) Sub(num1, num2 int) int {
	return num1 - num2
}

func t(a any) {
	v := reflect2.ValueOf(a).Elem()
	//t:=reflect2.TypeOf(a)
	k := v.Kind()
	if k != reflect2.Struct {
		fmt.Println("kind 错误")
		return
	}
	v.FieldByName("Num1").SetInt(100)
	v.FieldByName("Num2").SetInt(99)
	num := v.NumField()
	fmt.Println("STRUCT have", num)
	var f []reflect2.Value
	for i := 0; i < num; i++ {
		fmt.Printf("field[%v]=%v\n", i, v.Field(i))
		f = append(f, v.Field(i))
	}
	res := v.MethodByName("Sub").Call(f)
	fmt.Println("sub is ", res[0].Int())
}

func main() {
	m1 := M{}
	t(&m1)
}
