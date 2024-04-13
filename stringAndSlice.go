package main

import "fmt"

func main() {
	str := "hello,world"
	slice := str[6:] //string里面的值不可改变
	fmt.Println(slice)

	//要修改字符串的话要把string转换成[]byte或者[]rune
	arr := []byte(str)
	arr[0] = 'H'
	str = string(arr)
	fmt.Println(str)

	//Byte按字节编码，不能转换成中文
	//将string转换成[]rune就可以输出中文
	arr1 := []rune(str)
	arr1[5] = '中'
	str = string(arr1)
	fmt.Println(str)
}
