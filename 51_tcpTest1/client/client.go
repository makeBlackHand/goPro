package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:46666")
	defer con.Close()
	if err != nil {
		fmt.Println("client err", err)
		return
	}
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入【终端】
	for {
		liner, errs := reader.ReadString('\n')
		liner = strings.Trim(liner, " \r\n")
		if errs != nil {
			println("reader err", errs)
		}
		if liner == "exit" {
			fmt.Println("客户端退出")
			break
		}
		n, erro := con.Write([]byte(liner + "\n"))
		if erro != nil {
			fmt.Println("con.write err", erro)
		}
		fmt.Printf("客户端发送了%v个字节数据\n", n-1) //没算换行
		fmt.Println("client con 成功", con)
	}

}
