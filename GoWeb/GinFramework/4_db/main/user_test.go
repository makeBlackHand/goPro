package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

// TestMain函数可以在测试函数执行前做些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始:")
	//通过m.run来执行测试函数
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	//通过t.run()测试自测试程序
	t.Run("测试查询用户:", testUser_GetUsers)
	t.Run("测试查询用户:", testUser_GetUserById)
	t.Run("测试添加用户:", testUser_AddUser)
}

// 不以Test开头默认不执行,将他设成子测试程序
func testUser_AddUser(t *testing.T) {
	fmt.Println("子测试程序测试中:")
	//user := &User{}
	//user.AddUser()
	//user.AddUser2()
}

// 测试根据id获取一个user
func testUser_GetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录：")
	user := &User{
		Id: 1,
	}
	u, _ := user.GetUserById()
	fmt.Println("得到的user是:", u)
}

// 测试获取所有记录
func testUser_GetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录：")
	user := &User{}
	u, _ := user.GetUsers()
	for i, v := range u {
		fmt.Printf("得到第%d个user是:%v", i+1, v)
	}

}
