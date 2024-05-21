package process

import (
	"encoding/json"
	"fmt"
	"goPro/BasicSyntax/chatRoom/client/util"
	"goPro/BasicSyntax/chatRoom/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGroupMes jason err", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMes jason err", err)
		return
	}
	tf := &util.Transfer{
		Con: CurUser.Con,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendGroupMes writePkg err", err)
		return
	}
	return
}
