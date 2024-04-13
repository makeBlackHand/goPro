package main

import "fmt"

type Calcul struct {
	num1 int
	num2 int
}

func (c *Calcul) jia() int {
	return c.num2 + c.num1
}
func (c *Calcul) jian() int {
	return c.num1 - c.num2
}
func (c *Calcul) cheng() int {
	return c.num2 * c.num1
}
func (c *Calcul) chu() float64 {
	return float64(c.num1) / float64(c.num2)
}

func (c *Calcul) calcul(num1, num2 float64, zifu byte) (sum float64) {
	switch zifu {
	case '+':
		sum = num1 + num2
		return sum
	case '-':
		sum = num1 - num2
		return sum
	case '*':
		sum = num1 * num2
		return sum
	case '/':
		sum = num1 / num2
		return sum
	default:
		fmt.Println("输入有误")

	}
	return
}

func main() {
	num := Calcul{10, 20}
	sub := num.jia()
	substract := num.jian()
	cheng := num.cheng()
	chu := num.chu()
	fmt.Println(sub)
	fmt.Println(substract)
	fmt.Println(cheng)
	fmt.Println(float64(chu))

	sum := num.calcul(100, 50, '+')
	fmt.Println(sum)

}
