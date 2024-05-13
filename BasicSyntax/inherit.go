package main

import "fmt"

type Student0 struct {
	name  string
	age   int
	score float64
}

func (stu *Student0) inputScore(score float64) {
	stu.score = score
}
func (stu *Student0) showScore() {
	fmt.Printf("%v的成绩为:%v\n", stu.name, stu.score)
}

type Pupil struct {
	Student0
}

func (pupil *Pupil) test() {
	fmt.Println("小学生在测试。。。")
}

type Graduate struct {
	Student0
}

func (graduate *Graduate) test() {
	fmt.Println("大学生在测试。。。")
}

func main() {
	p := Pupil{}
	p.name = "张三"
	p.Student0.age = 99
	p.test()
	p.inputScore(49.6)
	p.Student0.showScore()

	g := Graduate{}
	g.Student0.name = "李四"
	g.Student0.age = 100
	g.test()
	g.inputScore(99.0)
	g.Student0.showScore()

}
