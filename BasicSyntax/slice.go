package main

import "fmt"

func main() {
	//第一种定义切片方法
	var arr = [...]int{1, 2, 3, 4, 5}
	slice := arr[1:3] //引用arr数组从下标1-3(不包括3)
	//slice在此次有三个地址第一个是保存引用arr的下标1的地址
	//第二个存放slice长度，第三个存放slice容量

	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) //一般容量为长度两倍
	fmt.Printf("arr[1]的地址:%p\n", &arr[1])
	fmt.Printf("slice[1]的地址:%p\n", &slice[0])

	//第二种定义切片方法make定义，切片中有三个定义第一个是数据类型，第二个是长度，第三个是容量（容量>=长度）
	var arr1 []int = make([]int, 6, 10) //只能通过切片修改
	fmt.Println(arr1)
	arr1[1] = 10
	arr1[4] = 20
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))

	//第三种定义一个切片，直接指定数组
	var slice1 []string = []string{"java", "vue", "c++"}
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice01 := slice[0:1] //切片可以在切片上切片
	fmt.Println(slice01)
	slice01[0] = 100
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(slice01)
	//切片上的切片改变数值 令其切片和数组都改变
}
