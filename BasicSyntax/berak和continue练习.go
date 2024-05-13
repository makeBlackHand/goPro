package main

import "fmt"

func main() {
	var time int
	var sum float64 = 100000
	for sum >= 1000 {
		if sum > 50000 {
			time++
			sum -= sum * 0.05
		} else if sum <= 50000 {
			time++
			sum -= 1000
		} else {
			break
		}
	}
	fmt.Printf("time=%d,sum=%f", time, sum)

	// var year myInt = 2019
	//ere := year.isRyear()
}

//type myInt int
//type myBool bool
//
//func (year myInt) isRyear() (res myBool) {
//	res = year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)
//	return
//}
