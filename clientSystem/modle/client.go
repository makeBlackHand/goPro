package modle

import "fmt"

type Client struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewClient2(name string, gender string, age int, phone string, email string) Client {
	return Client{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
func NewClient(id int, name string, gender string, age int, phone string, email string) Client {
	return Client{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (c Client) GetClient() string {
	str := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
		c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return str
}
