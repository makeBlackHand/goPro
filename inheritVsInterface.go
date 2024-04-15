package main

import "fmt"

type Monkey struct {
	Name string
}
type JiNeng interface {
	Fly()
	Swim()
}

func (monkey *Monkey) climb() {
	fmt.Println(monkey.Name, "会爬树")
}
func (monkey *Monkey) Fly() {
	fmt.Println(monkey.Name, "会飞翔")
}
func (monkey *Monkey) Swim() {
	fmt.Println(monkey.Name, "会游泳")
}

type LittleMonkey struct {
	Monkey
}

//

func main() {
	monkey := LittleMonkey{
		Monkey{Name: "斗战胜佛"},
	}
	monkey.climb()
	monkey.Fly()
	monkey.Swim()
}
