package main

import "fmt"

func main() {
	var arr [3][4]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	fmt.Println(arr)
}
