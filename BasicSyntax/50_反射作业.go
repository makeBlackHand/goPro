package main

import (
	"fmt"
	reflect2 "reflect"
)

type Cals struct {
	Num1 int
	Num2 int
}

func (c Cals) GetSub(num1, num2 int) int {
	return num1 - num2
}

func testing(a any) {
	//ty := reflect2.TypeOf(a)
	va := reflect2.ValueOf(a).Elem()
	k := va.Kind()
	if k != reflect2.Struct {
		println("struct err")
		return
	}
	va.FieldByName("Num1").SetInt(100)
	va.FieldByName("Num2").SetInt(99)
	num := va.NumField()
	var nums []reflect2.Value = make([]reflect2.Value, 2)

	for i := 0; i < num; i++ {
		fmt.Printf("field[%v]=%v\n", i, va.Field(i))
		fmt.Println(va.Field(i))
		//nums = append(nums, va.Field(i))
		nums[i] = va.Field(i)
	}
	fmt.Println("LEN", len(nums))
	res := va.MethodByName("GetSub").Call(nums)
	//va.Elem().MethodByName("Print").Call(nil)
	//fmt.Println(len(res))
	fmt.Println(res[0])
	fmt.Printf("%T %T", res[0], res)
}
func main() {
	cal := Cals{}
	testing(&cal)
}
