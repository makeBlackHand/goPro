package main

import (
	"fmt"
	"goPro/src/p33_aaaaa/model"
)

func main() {
	p := model.NewPerson("ass")
	if err := p.SetName("aaa"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.GetName())

}
