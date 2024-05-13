package modle

import (
	"fmt"
)

type FamilyAccount struct {
	num     int
	balance float64
	money   float64
	note    string
	state   string
	ok      bool
	detail  string
	flag    bool
	account string
	psw     int
}

func NewAccount() *FamilyAccount {
	return &FamilyAccount{
		balance: 1000.0,
		money:   0.0,
		note:    "",
		state:   "",
		ok:      false,
		detail:  "收支\t账户金额\t收支金额\t说明",
		flag:    true,
		account: "",
	}
}

func (familyAccount *FamilyAccount) Register() {
	for {
		fmt.Println("-----请登录家庭收支记账系统-----")
		fmt.Printf("请输入账号：")
		fmt.Scanln(&familyAccount.account)
		fmt.Printf("请输入密码：")
		fmt.Scanln(&familyAccount.psw)
		if familyAccount.psw == 123456 && familyAccount.account == "治黑手" {
			familyAccount.Menu()
			break
		} else if familyAccount.psw != 123456 && familyAccount.account != "治黑手" {
			fmt.Println("账号和密码错误")
		} else if familyAccount.psw != 123456 {
			fmt.Println("密码错误")
		} else if familyAccount.account != "治黑手" {
			fmt.Println("账号错误")
		}
	}
}
func (familyAccount *FamilyAccount) Menu() {
	fmt.Println("欢迎用户", familyAccount.account, "登录")
	for familyAccount.flag {
		fmt.Println("\n-------------家庭收支记账软件-------------")
		fmt.Println("		1.收支明细")
		fmt.Println("		2.登记收入")
		fmt.Println("		3.登记支出")
		fmt.Println("		4.转账")
		fmt.Println("		5.退出")
		fmt.Printf("         请输入（1-5）:")
		fmt.Scanln(&familyAccount.num)
		switch familyAccount.num {
		case 1:
			familyAccount.Record()
		case 2:
			familyAccount.Add()
		case 3:
			familyAccount.Delete()
		case 4:
			familyAccount.Transfer()
		case 5:
			familyAccount.Exit()
		default:
			fmt.Println("请输入正确选项!")
		}
	}
}
func (familyAccount *FamilyAccount) Record() {
	fmt.Println("------------当前收支明细记录------------")
	if familyAccount.ok {
		fmt.Println(familyAccount.detail)
	} else {
		println("还没有收支，快来一笔吧")
	}
}
func (familyAccount *FamilyAccount) Add() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&familyAccount.money)
	familyAccount.balance += familyAccount.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&familyAccount.note)
	familyAccount.detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t %v", familyAccount.balance, familyAccount.money, familyAccount.note)
	familyAccount.ok = true
}
func (familyAccount *FamilyAccount) Delete() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&familyAccount.money)
	if familyAccount.balance >= familyAccount.money {
		familyAccount.balance -= familyAccount.money
	} else {
		fmt.Println("余额不足")
		return
	}
	fmt.Println("本次支出说明:")
	fmt.Scanln(&familyAccount.note)
	familyAccount.detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t %v", familyAccount.balance, familyAccount.money, familyAccount.note)
	familyAccount.ok = true
}
func (familyAccount *FamilyAccount) Transfer() {
	var nam, or string
	var mon float64
	a := NewAccount()
	a.balance = 99.0
	for {
		fmt.Println("请输入转账的账户:")
		fmt.Scanln(&nam)
		a.account = nam
		fmt.Println("请输入转账金额:")
		fmt.Scanln(&mon)
		if mon > familyAccount.balance {
			fmt.Println("你没有这么多钱！！！")
		} else {
			for {
				familyAccount.balance -= mon
				a.balance += mon
				fmt.Println("转账成功")
				fmt.Println("你的余额为:", familyAccount.balance)
				fmt.Println(nam, "的余额为:", a.balance)
				fmt.Println("你还要继续转账吗！[y]：是 [n]")
				fmt.Scanln(&or)
				if or == "n" {
					break
				}
			}
			break
		}
	}
}
func (familyAccount *FamilyAccount) Exit() {
	println("\n你确定要退出吗？ [y]:是\t[n]:否")
	for {
		fmt.Scanln(&familyAccount.state)
		if familyAccount.state == "y" || familyAccount.state == "n" {
			break
		}
		println("输入错误,请重新输入")
	}
	if familyAccount.state == "y" {
		familyAccount.flag = false
	}
}
