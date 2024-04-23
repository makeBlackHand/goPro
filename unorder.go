package main

import (
	"encoding/json"
	"fmt"
)

type Monste struct {
	Name  string
	Age   int
	Birth string
	Sal   float64
	Skill string
}

func UnStruct() {
	str := "{\"Name\":\"牛魔王\",\"Age\":190,\"Birth\":\"2000-1-3\",\"Sal\":199.99,\"Skill\":\"还我飘飘拳\"}\n"
	var m Monste
	err := json.Unmarshal([]byte(str), &m) //必须要去取地址里面的值才改变
	if err != nil {
		fmt.Println("反序列化失败！！", err)
	}
	fmt.Printf("反序列化后monste=%v\n", m)
}

func UnMap() {
	str := "{\"Age\":\"333\",\"Name\":\"铁扇公主\",\"Skil\":\"呼风唤雨\"}\n"
	var a map[string]interface{} //反序列化不需要make分配空间自动分配空间
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("map=", a)
}

func UnSlice() {
	//str := "[{\"age\":\"18\",\"name\":\"jack\"},{\"address\":[\"墨西哥\",\"旧金山\"],\"age\":\"1800\",\"name\":\"杰克吗\"}]\n"
	str := TestSlice1()
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("map=", slice)
}

func TestSlice1() string {
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

func main() {
	UnStruct()
	UnMap()
	UnSlice()
}
