package main

import (
	"fmt"
	"goPro/BasicSyntax/chatRoom/client/process"
	"os"
)

var (
	Id  int
	Pwd string
)

func main() {
	var option int
	//var loop bool = true
	for {
		println("--------------欢迎登陆多人聊天室--------------")
		println("\t\t1.登陆聊天室")
		println("\t\t2.注册用户")
		println("\t\t3.退出系统")
		fmt.Printf("\t\t请选择（1~3）:")
		fmt.Scanf("%d\n", &option)
		switch option {
		case 1:
			println("登陆聊天室")
			fmt.Printf("请输入用户的id:")
			fmt.Scanf("%d\n", &Id)
			fmt.Printf("请输入用户的密码:")
			fmt.Scanf("%s\n", &Pwd)
			up := &process.UserProcess{}
			up.Login(Id, Pwd)
			//loop = false
		case 2:
			println("注册用户")
			//loop = false
		case 3:
			println("系统已退出。。。。")
			os.Exit(0)
			//loop = false
		default:
			println("输入错误,请重输")
		}
	}
	//if option == 1 {
	//
	//	//login(Id, Pwd)
	//	//if err != nil {
	//	//	fmt.Println("登录错误：", err)
	//	//} else {
	//	//	fmt.Println("登录成功")
	//	//}
	//} else if option == 2 {
	//
	//}

}
