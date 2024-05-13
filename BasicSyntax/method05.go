package main

import "fmt"

type Student02 struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (s *Student02) say() string {
	s1 := fmt.Sprintf("s的信息：name=%v,gender=%v,age=%v,id=%v,score=%v",
		s.name, s.gender, s.age, s.id, s.score)
	return s1
}

func main() {
	var s = Student02{
		"法外狂徒",
		"不伦不类",
		-101,
		180,
		90.5,
	}
	fmt.Println(s.say())

}
