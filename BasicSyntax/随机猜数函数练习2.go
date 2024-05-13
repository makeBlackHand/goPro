package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test1() (i int) {
	var n int
	for j := 0; j < 10; j++ {
		println("请输入猜的数字")
		fmt.Scan(&n)
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100) + 1
		i++
		if n != num {
			println("输入错误")
			println("正确答案=", num)
		}
		if j == 9 && n != num {
			i++
		}
	}
	return
}

func main() {
	num := test1()
	switch num {
	case 1:
		println("那你真是个天才")
	case 2, 3:
		println("你很聪明，快赶上我了")
	case 4, 5, 6, 7, 8, 9:
		println("一般")
	case 10:
		println("可算猜对了")
	default:
		println("说你傻呢")

	}

}
