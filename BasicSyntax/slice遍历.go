package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{1, 2, 3, 4, 5}
	slice := arr[:]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]=%d\n", i, slice[i])
	}

	for index, value := range slice {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}

}
