package main

import "fmt"

func main() {
	var num int
	var balance float64 = 10000.0
	var money float64
	var note string
	var state string
	var ok bool = false
	detail := "收支\t账户金额\t收支金额\t说明"
	flag := true
	for flag {
		fmt.Println("\n-------------家庭收支记账软件-------------")
		fmt.Println("		1.收支明细")
		fmt.Println("		2.登记收入")
		fmt.Println("		3.等级支出")
		fmt.Println("		4.退出")
		fmt.Printf("         请输入（1-4）:")
		fmt.Scanln(&num)
		switch num {
		case 1:
			fmt.Println("------------当前收支明细记录------------")
			if ok {
				fmt.Println(detail)
			} else {
				println("还没有收支，快来一笔吧")
			}

		case 2:
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t %v", balance, money, note)
			ok = true
		case 3:
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			if balance >= money {
				balance -= money
			} else {
				fmt.Println("余额不足")
				break
			}
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t %v", balance, money, note)
			ok = true
		case 4:
			println("\n你确定要退出吗？ [y]:是\t[n]:否")
			for {
				fmt.Scanln(&state)
				if state == "y" || state == "n" {
					break
				}
				println("输入错误,请重新输入")
			}
			if state == "y" {
				flag = false
			}

		default:
			fmt.Println("请输入正确选项!")

		}
	}
	fmt.Println("已退出家庭收支软件系统。。。")
}
