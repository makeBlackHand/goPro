package main

import (
	"fmt"
	reflect2 "reflect"
)

type Studentss struct {
	Name string
	Age  int
}

type Monsterr struct {
	Name string
	Age  int
}

func reflect(a interface{}) {
	rTyp := reflect2.TypeOf(a)
	fmt.Println(rTyp)
	fmt.Printf("rType is %v\n", rTyp)
	v := reflect2.ValueOf(a)
	sum := v.Int() + 10
	println(sum)
	fmt.Println(v)
	iv := v.Interface()
	num2 := iv.(int)
	println("num2:", num2)
}

func reflects(a interface{}) {
	rTyp := reflect2.TypeOf(a)
	fmt.Printf("rType is %v\n", rTyp)
	v := reflect2.ValueOf(a)
	fmt.Println(v)
	iv := v.Interface()
	num2, ok := iv.(Studentss) //不能直接取值要类型断言再取值
	if ok {
		println("num2.name=", num2.Name)
	}
	kind1 := rTyp.Kind()
	kind2 := v.Kind()
	fmt.Printf("kind1=%v, kind2=%v\n", kind1, kind2)

}

func main() {
	//var num int = 100
	//reflect(num)
	stu := Studentss{
		Name: "tom",
		Age:  19,
	}
	reflects(stu)
}
