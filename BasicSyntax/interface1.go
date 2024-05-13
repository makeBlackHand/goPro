package main

import "fmt"

type B2 interface {
	Hello()
}
type A2 interface {
	Say()
}
type Monster1 struct {
}

func (m Monster1) Say() {
	fmt.Println("monster say")
}
func (m Monster1) Hello() {
	fmt.Println("monster hello")
}
func main() {
	var m Monster1
	m.Hello()
	m.Say()

	var a float64 = 9.9
	var t interface{}
	t = a
	fmt.Println(t) //任何变量都可以赋值给空接口并输出

}
