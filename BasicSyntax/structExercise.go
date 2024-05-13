package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var f1 Person //第一种定义方法类似var a int
	f1.Age = 99
	f1.Name = "哈哈哈"
	fmt.Println(f1)

	f2 := Person{"张三", 15} //第二种定义方法直接定义
	fmt.Println(f2.Name)

	f3 := new(Person) //第三种定义方法用new原来是*f3指针设计者底层进行处理直接f3
	f3.Age = 18       //用法一样：(*f3).Age
	f3.Name = "李四"
	fmt.Println(f3)

	var f4 *Person = &Person{"mike", 66} //可以直接赋值也可以下面赋值
	f4.Age = 199
	f4.Name = "王五"
	fmt.Println(f4)

	p1 := Person{"mike", 10}
	var p2 *Person = &p1
	p2.Name = "lilian"
	fmt.Println(p2.Name)
}
