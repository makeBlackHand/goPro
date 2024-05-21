package process

import (
	"encoding/json"
	"fmt"
	"goPro/BasicSyntax/chatRoom/client/util"
	"goPro/BasicSyntax/chatRoom/common/message"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	var SmsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &SmsMes)
	if err != nil {
		fmt.Println("SendGroupMes 反序列化失败")
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes 序列化出错")
		return
	}
	//遍历服务端的onlineUser
	for id, up := range userMgr.onlineUser {
		//过滤一下不用发给自己
		if id == SmsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Con)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, con net.Conn) {
	tf := &util.Transfer{
		Con: con,
		Buf: [8096]byte{},
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败")
	}
}
