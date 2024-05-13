package main

import (
	"fmt"
	"io"
	"net"
)

func t(n net.Conn) {
	defer n.Close()
	for {
		b := make([]byte, 1024)
		n1, e1 := n.Read(b)
		if e1 == io.EOF {
			fmt.Println("客户端退出")
			return
		}
		fmt.Println(string(b[:n1]))
	}

}

func main() {
	fmt.Println("开始监听")
	c, e := net.Listen("tcp", "127.0.0.1:64444")
	if e != nil {
		fmt.Println("监听失败")
		return
	}
	defer c.Close()
	for {
		println("连接中")
		c1, e1 := c.Accept()
		if e1 != nil {
			fmt.Println("连接失败")
		} else {
			fmt.Println("con=", c1.RemoteAddr().String(), "客户端=", c1.RemoteAddr().String())
		}
		go t(c1)
	}
}
