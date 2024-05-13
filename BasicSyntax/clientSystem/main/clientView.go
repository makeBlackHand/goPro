package main

import (
	"fmt"
	"goPro/clientSystem/clientService"
	"goPro/clientSystem/modle"
	"strconv"
	"time"
)

type ClientView struct {
	Key          int
	Loop         bool
	clientServie *clientService.ClientService
}

func (c *ClientView) list() {
	client := c.clientServie.List()
	fmt.Println("-----------客户列表-----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(client); i++ {
		fmt.Println(client[i].GetClient())
	}
	fmt.Println("\n\n-----------客户列表完成-----------")
}

func (c *ClientView) Add() {
	startTime := time.Now()
	fmt.Println("------------添加客户------------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scan(&name)
	fmt.Print("性别：")
	gender := ""
	fmt.Scan(&gender)
	//var b bool = true
	//var age int
	//var b interface{}
	//fmt.Print("年龄：")
	//fmt.Scan(&b)
	//
	//b = age
	//_, v := b.(int)
	//if v != true {
	//	for {
	//		fmt.Println("请重新输入年龄:")
	//		fmt.Scan(&age)
	//		b = age
	//		_, v := b.(int)
	//		if v {
	//			break
	//		}
	//	}
	//}

a1:
	var ages string
	fmt.Print("年龄:")
	fmt.Scan(&ages)
	num, err := strconv.Atoi(ages)
	if err != nil {
		goto a1
	}
	//fmt.Printf("ages %T = %d", num, num)

	fmt.Print("电话：")
	phone := ""
	fmt.Scan(&phone)
	fmt.Print("email：")
	email := ""
	fmt.Scan(&email)
	client := modle.NewClient2(name, gender, num, phone, email)
	if c.clientServie.Add(client) {
		fmt.Println("--------添加成功--------\n\n")
	} else {
		fmt.Println("-----------添加失败-----------\n\n")
	}
	endTime := time.Now()
	fmt.Println("spend Time = ", endTime.Sub(startTime))

}

func (c *ClientView) Delete() {
	fmt.Println("------------删除客户------------")
	id := -1
	fmt.Print("\n请输入删除客户编号:")
	fmt.Scan(&id)
	if c.clientServie.Find(id) == -1 {
		fmt.Println("----------删除失败，输入id不对----------")
		return
	}

	var s string
	fmt.Print("请确认是否删除? [y]:是  [n]:否 = ")
	fmt.Scan(&s)
	if s == "y" || s == "Y" {
		flag := c.clientServie.Delete(id)
		if flag {
			fmt.Println("---------删除成功---------")
		} else {
			fmt.Println("----------删除失败，输入id不对----------")
		}
	}

}

func (c *ClientView) AlertMenu() {
	num := 0
	age := 0
	name, gender, phone, email := "", "", "", ""
	fmt.Println("----------修改客户----------")
	fmt.Print("请选择修改客户编号：")
	fmt.Scan(&num)
	fmt.Println("CHOOSED = ", num)
	if num > len(c.clientServie.Clients) {
		fmt.Println("下标越界")
		return
	}
	fmt.Printf("姓名(%v):", c.clientServie.Clients[num-1].Name)
	fmt.Scan(&name)
	fmt.Printf("性别(%v):", c.clientServie.Clients[num-1].Gender)
	fmt.Scan(&gender)
	fmt.Printf("年龄(%v):", c.clientServie.Clients[num-1].Age)
	fmt.Scan(&age)
	fmt.Printf("电话(%v):", c.clientServie.Clients[num-1].Phone)
	fmt.Scan(&phone)
	fmt.Printf("邮箱(%v):", c.clientServie.Clients[num-1].Email)
	fmt.Scan(&email)

	if c.clientServie.Alert(num, name, gender, age, phone, email) {
		fmt.Println("----------修改成功----------")
	} else {
		fmt.Println("--------------修改失败--------------")
	}

}

func (c *ClientView) Menu() {
	for c.Loop {
		fmt.Println("---------客户信息管理软件---------")
		fmt.Println("\t\t1.添加客户")
		fmt.Println("\t\t2.修改客户")
		fmt.Println("\t\t3.删除客户")
		fmt.Println("\t\t4.客户列表")
		fmt.Println("\t\t5.退出")
		fmt.Print("请选择(1-5):")
		fmt.Scan(&c.Key)
		switch c.Key {
		case 1:
			c.Add()
		case 2:
			c.AlertMenu()
		case 3:
			c.Delete()
		case 4:
			c.list()
		case 5:
			s := ""
			fmt.Print("请选择是否退出[y]:是 [n]:否")
			fmt.Scan(&s)
			if s == "y" || s == "Y" {
				c.Loop = false
			}
		default:
			fmt.Println("输入错误")
		}
	}
	fmt.Println("已经退出客户系统。。。")
}
func main() {
	client := ClientView{
		Loop: true,
	}
	client.clientServie = clientService.NewClientService()
	client.Menu()
}
