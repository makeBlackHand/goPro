package testExercise

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) store() bool {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败", err)
		return false
	}
	file := "E:/monster.ser"
	err1 := os.WriteFile(file, data, 123)
	if err != nil {
		fmt.Println("WriteFile错误", err1)
		return false
	}
	return true
}

func (m *Monster) restore() bool {
	file := "E:/monster.ser"
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("readfile 错误", err)
		return false
	}
	err1 := json.Unmarshal(data, m)
	if err1 != nil {
		fmt.Println("反序列化失败", err1)
		return false
	}
	return true
}
