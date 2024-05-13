package main

import "fmt"

type Account struct {
	name     string
	password int
	balance  float64
}

func (account *Account) deposit(money float64, password int) {
	if account.password != password {
		fmt.Println("密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("逗我玩呢！！")
		return
	}
	account.balance += money
	fmt.Println("成功存进：", account.balance)
}
func (account *Account) retrieve(money float64, password int) {
	if account.password != password {
		fmt.Println("密码不正确")
		return
	}
	if money <= 0 || money > account.balance {
		fmt.Println("逗我玩呢！！")
		return
	}
	account.balance -= money
	fmt.Println("成功取出：", money)
}
func (account *Account) query(password int) {
	if account.password != password {
		fmt.Println("密码不正确")
		return
	}

	fmt.Println("余额还有：", account.balance)
}
func main() {
	var num int
	var money float64
	var account Account
	account.password = 123
	account.balance = 123.5
	var password int
	for {
		println("请输入您要进行的功能：")
		println("[0]：退出服务")
		println("[1]：存款服务")
		println("[2]：取款服务")
		println("[3]：查询服务")
		fmt.Scan(&num)
		switch num {
		case 0:
			break
		case 1:
			fmt.Printf("请输入密码：")
			fmt.Scan(&password)
			if password == account.password {
				println("请输入存款金额")
				fmt.Scan(&money)
				account.deposit(money, 123)
			} else {
				fmt.Println("密码错误")
			}

		case 2:
			fmt.Printf("请输入密码：")
			fmt.Scan(&password)
			if password == account.password {
				println("请输入存款金额")
				fmt.Scan(&money)
				account.retrieve(money, 123)
			} else {
				fmt.Println("密码错误")
			}

		case 3:
			fmt.Printf("请输入密码：")
			fmt.Scan(&password)
			account.query(123)
		}
	}

}
