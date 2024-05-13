package main

import "fmt"

type Tiji struct {
	len   float64
	width float64
	high  float64
}

func (t Tiji) jisuan() float64 {
	sum := t.width * t.high * t.len
	return sum
}
func main() {

	var t Tiji
	t.len = 5
	t.width = 3
	t.high = 3.4
	sum := t.jisuan()
	fmt.Println(sum)
}
