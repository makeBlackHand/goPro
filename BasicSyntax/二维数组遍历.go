package main

import (
	"fmt"
)

func main() {
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v   ", arr[i][j])
		}
		fmt.Println()
	}

	for i, v := range arr {
		for i2, v2 := range v {
			fmt.Printf("arr[%v][%v]=%v   ", i, i2, v2)
		}

	}
}
