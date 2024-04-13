package main

import "fmt"

func main() {
	var arr [4][4]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	fmt.Println(arr)
	arr[0], arr[3] = arr[3], arr[0]
	arr[1], arr[2] = arr[2], arr[1]
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("arr[%d][%d]=%d  ", i, j, arr[i][j])
		}
		println()
	}
}
