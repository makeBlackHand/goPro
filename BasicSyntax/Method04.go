package main

import (
	"encoding/json"
	"fmt"
)

type Cal struct {
	Num1 int `json:"num1"`
	num2 int
}

func (c *Cal) cfb(num1 int) {
	for i := 1; i <= num1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d  ", i, j, i*j)
		}
		fmt.Println()
	}
}
func (c *Cal) daozhi() {
	var arr = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var arr2 [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//fmt.Printf("%d %d \n", i, j)

			arr2[i][j] = arr[j][i] //需要重新定义新数组存进去，原来数组不可重复用

			fmt.Printf("%d ", arr2[i][j])
		}
		println()
	}

}

func (c *Cal) daozhi2() {
	var arr = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var arr2 [3][3]int
	for i, ints := range arr {
		for i2, _ := range ints {
			arr2[i][i2] = arr[i2][i]
			fmt.Printf("%d\t", arr2[i][i2])
		}
		fmt.Println()
	}

}

func main() {
	var c Cal
	c.cfb(8)
	c.daozhi()
	s, _ := json.Marshal(c)
	fmt.Println(string(s))
	c.daozhi2()
}
