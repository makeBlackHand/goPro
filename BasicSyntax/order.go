package main

import (
	"encoding/json"
	"fmt"
)

type Monst struct {
	Name  string
	Age   int
	Birth string
	Sal   float64
	Skill string
}

func TestStruct() {
	monster := Monst{
		Name:  "牛魔王",
		Age:   190,
		Birth: "2000-1-3",
		Sal:   199.99,
		Skill: "还我飘飘拳",
	}
	data, err := json.Marshal(monster) //因为struct是结构体值类型
	if err != nil {
		println("序列化失败 err=", err)
		return
	}
	fmt.Printf("struct序列化后是%v\n", string(data))

}

func TestMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["Name"] = "铁扇公主"
	a["Age"] = "333"
	a["Skil"] = "呼风唤雨"
	data, err := json.Marshal(a) //map是引用类型
	if err != nil {
		println("序列化失败 err=", err)
		return
	}
	fmt.Printf("map序列化后是%v\n", string(data))
}
func TestSlice() string {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "18"
	slice = append(slice, m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "杰克吗"
	m2["age"] = "1800"
	m2["address"] = [2]string{"墨西哥", "旧金山"}
	slice = append(slice, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		println("序列化失败 err=", err)
	}
	//fmt.Printf("slice序列化后是%v\n", string(data))
	return string(data)
}

func TestFloat() { //基本类型序列化
	var f float64 = 44.567
	data, err := json.Marshal(f) //map是引用类型
	if err != nil {
		println("序列化失败 err=", err)
		return
	}
	fmt.Printf("map序列化后是%v\n", string(data))
}
func main() {
	TestStruct()
	TestMap()
	TestSlice()
	TestFloat() //对基本类型序列化=转成字符串
}
