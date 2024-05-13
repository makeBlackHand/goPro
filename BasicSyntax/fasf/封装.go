package fasf

import "fmt"

type person struct {
	Name    string
	age     int
	balance float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}
func (p *person) SetAge(age int) {
	if age > 0 && age <= 100 {
		p.age = age
	} else {
		fmt.Println("你是王八吗")
	}
}
func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetBalance(balance float64) {
	if balance >= 3000 && balance <= 10000 {
		p.balance = balance
	} else {
		fmt.Println("你是社畜吗")
	}
}
func (p *person) GetBalance() float64 {
	return p.balance
}

//func main() {
//
//}
