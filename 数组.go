package main

import (
	"fmt"
)

func main() {
	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 2.5
	hens[2] = 4.0
	hens[3] = 12.9
	hens[4] = 12.0
	hens[5] = 1.0
	var sum float64
	for i := 0; i < len(hens); i++ {
		sum += hens[i]
	}
	fmt.Printf("%.2f\n", sum) //输出小数点后两位
	fmt.Printf("%.2f\n", sum/6.0)

	var arr [6]int
	arr[1] = 2
	arr[5] = 1
	fmt.Println("arr=", arr)        //输出整个数组
	fmt.Printf("arr 地址：%p\n", &arr) //%p代表数组
	fmt.Printf("arr1 地址：%p\n", &arr[1])

	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 [3]string = [3]string{1: "tom", 0: "jack", 2: "hello"} //字符型数组可以定义指定位置
	fmt.Println("arr1=", arr1)
	fmt.Println("arr2=", arr2)

	//四种初始化数组方式
	var array1 [3]int = [3]int{1, 2, 3}
	fmt.Println(array1)

	var array2 = [3]int{1, 2, 3}
	fmt.Println(array2)

	var array3 = [...]int{1, 2, 3}
	fmt.Println(array3)

	var array4 = [...]int{1: 1, 0: 2, 2: 3}
	fmt.Println(array4)

	array5 := [...]string{1: "tom", 0: "jack", 2: "hello"}
	fmt.Printf("array5 type is %T", array5)
	fmt.Println(array5)
}
