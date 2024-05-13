package fea

import (
	"fmt"
	"goPro/fasf"
)

func main1() {
	a := fasf.NewPerson1("12", 12) //返回的是一个指针（地址）
	fmt.Println(*a)
	fmt.Println(a.GetAge())
}
