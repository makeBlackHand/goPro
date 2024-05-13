package main

import "fmt"

type Student1 struct {
	name string
	age  int
}

func (s Student1) judge(num int) {
	if num%2 != 0 {
		fmt.Println("是奇数")
	} else {
		fmt.Println("是偶数")
	}
}
func (s Student1) array(num1, num2, num int) {
	for i := 0; i < num1; i++ {
		for i := 0; i < num2; i++ {
			fmt.Printf("%d  ", num)
		}
		fmt.Println()
	}
}

func main() {
	num := Student1{"asd", 12}
	num.judge(12)
	num.judge(13)
	num.array(2, 3, 9)
}
