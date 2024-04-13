package main

import "fmt"

func main() {
	var monster []map[string]string
	monster = make([]map[string]string, 2) //先声明切片并分配空间
	if monster[0] == nil {
		monster[0] = make(map[string]string, 2) //再声明map分配空间
		monster[0]["name"] = "诸葛亮"
		monster[0]["Sex"] = "男"
	}
	if monster[1] == nil {
		monster[1] = make(map[string]string, 2)
		monster[1]["name"] = "刘备"
		monster[1]["Sex"] = "女"
	}
	newhero := map[string]string{
		"name": "吕奉先",
		"sex":  "人缘",
	}
	monster = append(monster, newhero)
	fmt.Println(monster)
}
