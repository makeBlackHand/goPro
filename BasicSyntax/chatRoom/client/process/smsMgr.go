package process

import (
	"encoding/json"
	"fmt"
	"goPro/BasicSyntax/chatRoom/common/message"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("smsMgr output 反序列化错误")
		return
	}

	info := fmt.Sprintf("用户%d:\t对大家说\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
