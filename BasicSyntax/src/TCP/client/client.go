package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	ipaddr = "localhost"
	port   = 8888
	proto  = "tcp"
)

func main() {
	var err error
	dialer, err := net.Dial(proto, fmt.Sprintf("%s:%d", ipaddr, port))
	if err != nil {
		println(err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			println(err)
		}
		dialer.Write([]byte(line))
		fmt.Print("> ")
	}
	defer dialer.Close()
}
