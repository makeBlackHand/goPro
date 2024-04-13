package main

import "fmt"

var a int = 9

func main() {
	for i := 1; i <= a; i++ {
		for k := 1; k <= a-i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= i*2-1; j++ {
			if j == 1 || j == i*2-1 || i == a {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}

		}
		println()
	}
}
