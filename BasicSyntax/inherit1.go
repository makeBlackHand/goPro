package main

import "fmt"

type A1 struct {
	Name string
	Age  int
}

func (a *A1) sayOK() {
	fmt.Println("A OK", a.Name)
}
func (a *A1) sayHello() {
	fmt.Println("A hello", a.Name)
}

type B1 struct {
	A1
	Name string
}

func (b *B1) sayOK() {
	fmt.Println("A OK", b.Name)
}

func main() {
	var b B1
	b.Name = "张三"
	b.sayOK()
	b.A1.Name = "zhangsan" //B1结构体没有这个方法会继承A1里的方法，但没赋值
	b.sayHello()
}
