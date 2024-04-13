package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c *Circle) area1() float64 {
	c.radius = 10
	return 3.14 * c.radius * c.radius
}

func main() {
	var c Circle
	c.radius = 2
	area := c.area()
	fmt.Println("面积是：", area)
	area1 := c.area1() //原来是(&c).area()这种写法，底层优化可以写成前面方式
	fmt.Println("面积是：", area1)
}
