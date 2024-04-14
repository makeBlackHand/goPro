package main

import (
	"fmt"
)

func TypeJudge(item ...interface{}) {
	for index, value := range item {
		switch value.(type) {
		case bool:
			fmt.Printf("第%v的类型是bool,值是%v\n", index, value)
		case int:
			fmt.Printf("第%v的类型是int,值是%v\n", index, value)
		case float64:
			fmt.Printf("第%v的类型是float64,值是%v\n", index, value)
		case string:
			fmt.Printf("第%v的类型是string,值是%v\n", index, value)
		default:
			fmt.Println("类型不匹配")
		case heros:
			fmt.Printf("第%v的类型是heros,值是%v\n", index, value)
		case *heros:
			fmt.Printf("第%v的类型是*heros,值是%v\n", index, value)
		}
	}
}

type heros struct {
}

func main() {
	i := 12
	f := 12.2
	b := true
	s := "南京"
	h := heros{}
	h1 := &heros{}
	TypeJudge(i, f, b, s, h, h1)
}
