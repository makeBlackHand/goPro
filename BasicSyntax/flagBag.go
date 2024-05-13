package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名默认为空")
	//&user用来接收命令行中-u后面的参数值
	//"u"，就是-u的指定参数
	//""，默认值
	//"用户名默认为空"，说明
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号默认为3306") //记得要输入数字！
	flag.Parse()
	fmt.Printf("user=%v\t,pwd=%v\t,host=%v\t,port=%v\n",
		user, pwd, host, port)

}
