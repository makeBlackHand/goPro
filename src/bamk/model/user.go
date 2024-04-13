package model

import "fmt"

type user struct {
	accountno string
	balance   float64
	password  string
}

func NewUser(account string, password string, balance float64) *user {
	if len(account) < 5 || len(account) > 10 {
		fmt.Println("账号错误")
		return nil
	}
	if len(password) != 6 {
		fmt.Println("密码错误")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额错误")
		return nil
	}
	return &user{
		accountno: account,
		balance:   balance,
		password:  password,
	}
}
func (u *user) SetAccountNo(accountNo string) {
	if len(accountNo) < 5 || len(accountNo) > 10 {
		fmt.Println("账号错误")
	} else {
		u.accountno = accountNo
	}
}
func (u *user) SetPassword(password string) {
	if len(password) != 6 {
		fmt.Println("密码错误")
	} else {
		u.password = password
	}
}
func (u *user) SetBalance(balance float64) {
	if balance < 20 {
		fmt.Println("余额错误")
	} else {
		u.balance = balance
	}
}
func (u *user) GetAccountNo() string {
	return u.accountno
}
func (u *user) GetPassword() string {
	return u.password
}

func (u *user) GetBalance() float64 {
	return u.balance
}
