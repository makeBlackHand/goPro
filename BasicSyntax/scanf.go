package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入身高:")
	fmt.Scanln(&a)
	fmt.Printf("身高:%d", a)
	//var a float64 = 3.0
	//var b float64 = 100.0
	//var c float64 = 6.0
	//m := b*b - 4*a*c
	//if m > 0 {
	//	x1 := (-b + math.Sqrt(m)) / 2 * a
	//	x2 := (-b - math.Sqrt(m)) / 2 * a
	//	fmt.Printf("x1=%f,x2=%f\n", x1, x2)
	//} else if m < 0 {
	//	x1 := (-b + math.Sqrt(m)) / 2 * a
	//	fmt.Printf("x1=%f", x1)
	//} else {
	//	fmt.Println("无解")
	//}
}
