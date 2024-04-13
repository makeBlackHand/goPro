package main

import "fmt"

func main() {
	var slice []int = []int{1, 2, 3, 4, 5}
	slice = append(slice, 6, 7, 8) //slice本身没变化，要重新赋值给slice
	fmt.Println(slice)
	slice = append(slice, slice...) //可以在切片后追加一个切片，追加切片时要加...
	fmt.Println(slice)

	var a []int = []int{100, 200, 300}
	var b []int = make([]int, 10, 100)
	copy(b, a) //把a切片拷贝到b 其余默认是0
	//拷贝时必须都是切片
	//拷贝时拷贝五个数，但只能放一个数，取第一个数复制到里
	fmt.Println(b)
}
