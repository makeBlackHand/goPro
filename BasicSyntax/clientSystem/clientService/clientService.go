package clientService

import (
	"fmt"
	"goPro/clientSystem/modle"
)

type ClientService struct {
	Clients   []modle.Client //用来存放多少个客户
	ClientNum int            //存放客户数量
	file1     modle.File
}

func NewClientService() *ClientService {
	clientService := &ClientService{}
	clientService.ClientNum = 1
	client := modle.NewClient(1, "张三", "男", 18, "110", "qq")
	clientService.Clients = append(clientService.Clients, client)
	return clientService
}

func (c *ClientService) List() []modle.Client {
	return c.Clients
}

func (c *ClientService) Add(client modle.Client) bool {
	c.ClientNum++
	client.Id = c.ClientNum
	c.Clients = append(c.Clients, client)
	return true
}

func (c *ClientService) Delete(id int) bool {
	index := c.Find(id)
	if index == -1 {
		return false
	}
	c.Clients = append(c.Clients[:index], c.Clients[index+1:]...)
	return true
}

func (c *ClientService) Alert(id int, name string, gender string, age int, phone string, email string) bool {
	fmt.Println("len = ", len(c.Clients))

	index := c.Find(id)
	fmt.Println("custome = ", c.Clients[index])
	if index == -1 {
		return false
	}
	fmt.Println("开始修改：")
	if name != "" {
		c.Clients[index].Name = name
		c.file1.Store("客户Name修改成功\t\n")
	}
	if gender != "" {
		c.Clients[index].Gender = gender
		c.file1.Store("客户Gender修改成功\t\n")
	}
	if age != 0 {
		c.Clients[index].Age = age
	}
	if phone != "" {
		c.Clients[index].Phone = phone
	}
	if email != "" {
		c.Clients[index].Email = email
	}

	//if name == "\n" {
	//	c.Clients[index].Age = age
	//	c.Clients[index].Gender = gender
	//	c.Clients[index].Phone = phone
	//	c.Clients[index].Email = email
	//} else if gender == "\n" {
	//	c.Clients[index].Name = name
	//	c.Clients[index].Age = age
	//	c.Clients[index].Phone = phone
	//	c.Clients[index].Email = email
	//} else if value, _ := strconv.Atoi("\n"); value == age {
	//	c.Clients[index].Name = name
	//	c.Clients[index].Gender = gender
	//	c.Clients[index].Phone = phone
	//	c.Clients[index].Email = email
	//} else if phone == "\n" {
	//	c.Clients[index].Name = name
	//	c.Clients[index].Age = age
	//	c.Clients[index].Gender = gender
	//	c.Clients[index].Email = email
	//} else if email == "\n" {
	//	c.Clients[index].Name = name
	//	c.Clients[index].Age = age
	//	c.Clients[index].Gender = gender
	//	c.Clients[index].Phone = phone
	//}
	c.file1.Store("客户信息修改成功\t\n")
	return true

}

func (c *ClientService) Find(num int) int {
	index := -1
	for i := 0; i < len(c.Clients); i++ {
		if c.Clients[i].Id == num {
			index = i
		}
	}
	return index
}
