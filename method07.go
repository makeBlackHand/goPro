package main

import "fmt"

type Person01 struct {
	Name string
	Age  int
}

func (p Person01) ticket() {
	if p.Age > 18 {
		fmt.Printf("%v的年龄是%v超过了18，收费20\n", p.Name, p.Age)
	} else {
		fmt.Printf("%v的年龄是%v不到18，免门票\n", p.Name, p.Age)
	}

}
func main() {
	var p Person01
	for {
		println("请输入名字：")
		fmt.Scanln(&p.Name)
		if p.Name == "exit" {
			fmt.Println("已退出。。。")
			break
		}
		println("请输入年龄：")
		fmt.Scanln(&p.Age)
		p.ticket()
	}
}
