package main

import "fmt"

func main() {
	num1 := 10
	fmt.Printf("num1 type is %T,num1=%d,num1地址=%v\n", num1, num1, &num1)
	num2 := new(int) //*int 里面的地址默认是0的地址  num2再赋值时里面值改变地址不变
	fmt.Printf("num2 type is %T,num2=%d,num2地址=%v，num2的值为%v\n", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Printf("num2 type is %T,num2=%d,num2地址=%v，num2的值为%v\n", num2, num2, &num2, *num2)
	//new 给基本类型分配内存
	//make 给引用类型分配内存
}
