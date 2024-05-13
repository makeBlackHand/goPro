package process

import (
	"fmt"
	"goPro/BasicSyntax/chatRoom/client/util"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("--------xxx登陆成功-------")
	fmt.Println("--------1.用户在线列表-----")
	fmt.Println("--------2.发送消息--------")
	fmt.Println("--------3.信息列表--------")
	fmt.Println("--------4.退出系统--------")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示用户在线列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("退出系统中....")
		os.Exit(0)
	default:
		fmt.Println("输入错误，请重输！")

	}
}

// 和服务器保持通讯
func ServerProcessMes(con net.Conn) {
	tf := &util.Transfer{
		Con: con,
		Buf: [8096]byte{},
	}
	for {
		fmt.Println("客户端在等待服务端发送消息")
		mes, e := tf.ReadPkg()
		if e != nil {
			fmt.Println("客户端读取服务端失败")
			return
		}
		fmt.Println("mes=", mes)
	}
}
