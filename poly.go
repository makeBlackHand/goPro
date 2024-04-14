package main

import (
	"fmt"
)

type Usbs interface {
	Start()
	Stop()
}
type Phones struct {
	Name string
}

func (p Phones) Start() {
	fmt.Println("手机开始工作")
} //要实现接口里所有方法才叫实现这个接口
func (p Phones) Stop() {
	fmt.Println("手机停止工作")
}

type Cameras struct {
	Name string
}

func (c Cameras) Start() {
	fmt.Println("相机开始工作")
}
func (c Cameras) Stop() {
	fmt.Println("相机停止工作")
}
func main() {
	var usbArr [2]Usbs
	usbArr[0] = Phones{"小米"}
	usbArr[1] = Cameras{"尼康"}
	fmt.Println(usbArr)
}
