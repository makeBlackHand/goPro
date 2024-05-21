package main

import (
	"fmt"
	"goPro/BasicSyntax/chatRoom/common/message"
	process2 "goPro/BasicSyntax/chatRoom/server/process" //包名和方法名重复会报错
	"goPro/BasicSyntax/chatRoom/server/util"
	"io"
	"log"
	"net"
)

type Processor struct {
	Con net.Conn
}

// 根据传的数据不同，决定调用哪个函数
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	log.Println("处理所有消息")
	//根据客户端发送的消息种类不同来处理
	switch mes.Type {
	case message.LoginMesType: //处理登录逻辑
		up := &process2.UserProcess{
			Con: this.Con,
		}
		err = up.ServerProcessLogin(mes)
	case message.ResigterMesType: //处理注册逻辑
		up := &process2.UserProcess{
			Con: this.Con,
		}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型存在，无法处理。。")
	}
	return err
}

func (this *Processor) Process2() (err error) {
	log.Println("Process2")
	//循环客户端发送的消息
	for {
		tf := &util.Transfer{
			Con: this.Con,
			Buf: [8096]byte{},
		}
		log.Println("读取新消息前")
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了")
				return err
			} else {
				fmt.Println("readPkg err")
				return err
			}
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
