package main

import "fmt"

func main() {
	var a byte
	fmt.Scanf("%c", &a)
	switch a {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	default:
		fmt.Println("other")
	}

}
