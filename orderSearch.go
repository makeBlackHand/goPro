package main

import "fmt"

func order(name *[3]string, n string) (inex int) {
	inex = -1
	for i := 0; i < len(*name); i++ {
		if name[i] == n {
			inex = i
			fmt.Println("找到了", name[i], "下标为:", i)
			break
		} else {
			fmt.Println("没找到", n)
			break
		}
	}
	return
}
func main() {
	var n string
	var name = [3]string{"张三", "李四", "王五"}
	fmt.Scan(&n)
	order(&name, n)
}
