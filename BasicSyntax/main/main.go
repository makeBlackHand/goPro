package main

import (
	"fmt"
	"goPro/fasf"
)

func main() {
	p := fasf.NewPerson("小明")
	fmt.Println(p)
	p.SetAge(18)
	p.SetBalance(1800)
	fmt.Println(p)
}
