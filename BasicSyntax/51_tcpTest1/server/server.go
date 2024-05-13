package main

import (
	"fmt"
	"io"
	"net"
)

func process(con net.Conn) {
	defer con.Close()
	for {
		buf := make([]byte, 1024)
		//fmt.Println("服务器等待输入", con.RemoteAddr().String())
		n, err := con.Read(buf)
		//等待客户端通过con发送信息
		//如果客户端一直不write那协程一直阻塞
		if err == io.EOF {
			fmt.Println("客户端退出")
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "127.0.0.1:46666")
	if err != nil {
		fmt.Println("监听失败")
		return
	}
	defer listen.Close()

	for {
		println("等待客户端连接")
		con, erro := listen.Accept()
		if erro != nil {
			println("connect err", erro)
		} else {
			fmt.Println("con=", con, "客户端=", con.RemoteAddr().String())
		}
		go process(con)
	}
	//fmt.Println(listen)
}
