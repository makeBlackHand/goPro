package fasf

type person1 struct {
	Name string
	age  int
}

func NewPerson1(n string, a int) *person {
	return &person{
		Name: n,
		age:  a,
	}
} //让结构体首字母小写也可以调用

func (p *person) GetAge1() int {
	return p.age
} //让结构体属性首字母小写也可以调用

func main() {

}
