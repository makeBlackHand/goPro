package main

import "fmt"

func main() {
	var arr = [8]float64{1, 2, 3, 4, 5, 6, 7, 12}
	var sum float64
	var da, xiao, he int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	average := sum / float64(len(arr))
	fmt.Println("平均数为：", average)
	for i := 0; i < len(arr); i++ {
		if arr[i] > average {
			da++
		} else if arr[i] == average {
			he++
		} else {
			xiao++
		}
	}
	fmt.Printf("大于的有：%d，小于的有:%d,一样的有：%d", da, xiao, he)
}
