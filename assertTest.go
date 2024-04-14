package main

import (
	"fmt"
)

type Usb2 interface {
	Start()
	Stop()
}
type Phone1 struct {
	Name string
}

func (p Phone1) Start() {
	fmt.Println("手机开始工作")
} //要实现接口里所有方法才叫实现这个接口
func (p Phone1) Stop() {
	fmt.Println("手机停止工作")
}
func (p Phone1) Call() {
	fmt.Println(p.Name, "手机打电话")
}

type Camera1 struct {
	Name string
}

func (c Camera1) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera1) Stop() {
	fmt.Println("相机停止工作")
}

type Computers struct {
}

func (c Computers) Testing(u Usb2) {
	u.Start()
	if Phones, flag := u.(Phone1); flag {
		Phones.Call()
	}
	u.Stop()
}
func main() {
	var usbArr [2]Usb2
	usbArr[0] = Phone1{"小米"}
	usbArr[1] = Camera1{"尼康"}
	fmt.Println(usbArr)

	var c Computers
	for _, usb2 := range usbArr {
		c.Testing(usb2)
	}
}
