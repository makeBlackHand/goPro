package main

import (
	"fmt"
	"goPro/src/bamk/model"
)

func main() {
	p := model.NewUser("123456", "123456", 30)
	if p != nil {
		fmt.Println(p)
	} else {
		fmt.Println("创建失败")
	}
	p.SetBalance(700)
	p.SetAccountNo("456789")
	p.SetPassword("123456")
	fmt.Println(p)
}
