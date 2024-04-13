package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var sum int
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		if n == 99 {
			fmt.Println("sum=", sum)
			break
		} else {
			sum++
			fmt.Println("n=", n)
		}
	}

}
