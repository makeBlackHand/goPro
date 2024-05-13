package main

import "fmt"

type inter int
type Student struct {
	name string
	age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=%v,age=%v", stu.name, stu.age)
	return str
}

func (i inter) print() {
	println("i=", i)
}
func (i *inter) alter() {
	*i += 10
}

func main() {
	var i inter
	i = 11
	i.print()
	i.alter()
	fmt.Println(i)

	stu := Student{"马楼", 1}
	fmt.Println(&stu) //如果实现了String方法，系统会默认调用String方法

}
