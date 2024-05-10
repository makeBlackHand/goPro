package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	c, e := net.Dial("tcp", "127.0.0.1:64444")
	if e != nil {
		fmt.Println("client 连接失败")
		return
	}
	defer c.Close()
	read := bufio.NewReader(os.Stdin)
	for {
		r, er := read.ReadString('\n')
		r = strings.Trim(r, "\n")
		if er != nil {
			println("client read err")
		}
		if r == "exit" {
			fmt.Println("客户端退出")
			return
		}
		n1, err := c.Write([]byte(r))
		if err != nil {
			fmt.Println("wirte err")
			return
		}
		fmt.Println("写入了", n1, "个")
	}
}
