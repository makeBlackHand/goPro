package main

import "fmt"

func modify(user map[string]map[string]string, name string) {
	if user[name] != nil {
		user[name]["password"] = "123456"
	} else {
		user[name] = make(map[string]string, 2) //要先分配空间再写值
		user[name]["nickname"] = "nickname:" + name
		user[name]["password"] = "123456"
	}
}
func main() {
	var user map[string]map[string]string = make(map[string]map[string]string, 10)
	user["jordan"] = make(map[string]string, 2)
	user["jordan"]["nickname"] = "666"
	user["jordan"]["password"] = "888"
	modify(user, "jordan")
	modify(user, "mike")
	fmt.Println(user)
}
