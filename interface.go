package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Usb1 interface {
	Say()
}
type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
} //要实现接口里所有方法才叫实现这个接口
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (c *Computer) work(usb Usb) {
	usb.Start()
	usb.Stop()
}

type inter1 int

func (i inter1) Say() {
	fmt.Println("say i=", i)
}
func main() {
	p := Phone{}
	c := Camera{}
	c1 := Computer{}

	c1.work(p)
	c1.work(c)
	var i inter1 = 10
	var u Usb1 = i
	u.Say()
}
