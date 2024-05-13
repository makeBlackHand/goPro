package main

import "fmt"

type Person1 struct {
	name string
	age  int
}

func (p Person1) test() {
	fmt.Println("test p.name=", p.name)
}
func (p Person1) speak() {
	fmt.Printf("%s是好人", p.name)
}
func (p Person1) calcul() {
	var sum int
	for i := 0; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("1+2+3...+1000=", sum)
}
func (p Person1) calcul2(num int) {
	var sum int
	for i := 0; i <= num; i++ {
		sum += i
	}
	fmt.Printf("1+2+3...+num=%d", sum)
}
func (p Person1) calcul3(num, num2 int) int {
	sum := num + num2
	return sum
}
func main() {
	p := Person1{"莉莉安", 1}
	p.test() //test并不是函数是和Person1绑定的
	//方法是值传递
	p.speak()
	p.calcul()
	p.calcul2(5)
	sum := p.calcul3(10, 20)
	fmt.Println()
	fmt.Println(sum)
}
