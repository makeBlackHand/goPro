package main

import "fmt"

func arrTest(arr [3]int) {
	arr[0] = 10
}
func arrTest1(arr *[3]int) {
	arr[0] = 10
}

func main() {
	var str = [3]string{"张三", "李四", "王五"}
	for index, value := range str { //index是下表，value是值
		fmt.Printf("index=%v,value=%v\n", index, value)
	}
	var arr = [3]int{1, 2, 3}
	arrTest(arr) //main栈中的数组复制到arrTest栈main中arr不改变
	fmt.Println(arr)
	arrTest1(&arr)
	fmt.Println(arr)
}
