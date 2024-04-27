package main

import (
	"fmt"
	"net"
	"reflect"
)

type NetConn struct {
	Protocol string `json:"protocol"`
	IpAddr   string `json:"ip_addr"`
	Port     int    `json:"port"`
}

func (conn *NetConn) StartListener() {
	listener, err := net.Listen(conn.Protocol, fmt.Sprintf("%s:%d", conn.IpAddr, conn.Port))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("服务器已启动 " + fmt.Sprintf("%s:%d", conn.IpAddr, conn.Port))
	}
	defer listener.Close()
	for {
		connec, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Waiting for connection...")
		}
		go Process(connec)
	}

}

func Process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	netconn := NetConn{}
	reflectNet(&netconn)
	fmt.Println("main: ", netconn)
	netconn.StartListener()
}

func reflectNet(b interface{}) {
	rTpy := reflect.TypeOf(b)
	fmt.Println("rTpy =", rTpy)
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal =", rVal)

	fieldNums := rTpy.Elem().NumField()
	fmt.Println("fields:", fieldNums)
	for i := 0; i < fieldNums; i++ {
		tagName := rTpy.Elem().Field(i).Tag.Get("json")
		if tagName == "" {
			fmt.Println("tagName is empty")
		} else {
			fmt.Printf("字段[%d] = %v\n", i, tagName)
		}
	}

}
