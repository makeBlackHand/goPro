package main

import (
	"fmt"
	reflect2 "reflect"
)

type monsters struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string
}

func (m monsters) Print() {
	println("print start")
	fmt.Println(m)
	println("print end")
}

func (m monsters) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m monsters) Set(n string, n2 int, n3 string) {
	m.Name = n
	m.Age = n2
	m.Sex = n3
}

func TestStruct111(a interface{}) {
	va := reflect2.ValueOf(a)
	ty := reflect2.TypeOf(a)
	k := va.Kind()
	if k != reflect2.Ptr && va.Elem().Kind() == reflect2.Struct {
		println("not struct")
		return
	}
	num := va.Elem().NumField()
	fmt.Println("struct have ", num)

	va.Elem().FieldByName("Name").SetString("uuuu")
	//va.Elem().Field(0).SetString("凹凸曼")
	for i := 0; i < num; i++ {
		fmt.Printf("field[%v]=%v\n", i, va.Elem().Field(i))
		tagVal := ty.Elem().Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("i=%v tag=%v\n", i, tagVal)
		}
	}
	numMethod := va.Elem().NumMethod()
	println("方法有:", numMethod)

	//va.Elem().Method(1).Call(nil) //方法排序是按照ascll首字母排序的
	va.Elem().MethodByName("Print").Call(nil)
	//获取到第二个方法并传零个参数并调用
	var p []reflect2.Value
	p = append(p, reflect2.ValueOf(10))
	p = append(p, reflect2.ValueOf(20))
	res := va.Elem().Method(0).Call(p) //传去切片返回也是切片
	println("res=", res[0].Int())
}

func main() {
	var m = monsters{
		Name: "滴妈问",
		Age:  20,
		Sex:  "女",
	}
	TestStruct111(&m)
}
