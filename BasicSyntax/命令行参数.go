package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数有：", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, arg)
	}
}
