package main

import (
	"fmt"
	"strings"
)

func AddUpper() func(int) int {
	var n = 10
	var str = "hello"
	return func(i int) int {
		n += i
		str += "a"
		fmt.Println("str=", str)
		return n
	}
}
func makeSuffic(suffic string) func(string) string {
	return func(s string) string {
		if strings.HasSuffix(s, suffic) == false {
			return s + suffic
		}
		return s
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(19))

	f2 := makeSuffic(".jpg")
	fmt.Println(f2("hello"))
	fmt.Println(f2("bird.jpg"))
	fmt.Println(f2("apple.png"))
}
