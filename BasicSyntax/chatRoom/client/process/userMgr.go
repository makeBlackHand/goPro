package process

import (
	"fmt"
	"goPro/BasicSyntax/chatRoom/client/model"
	"goPro/BasicSyntax/chatRoom/common/message"
)

// 客户端的map
var OnlineUser map[int]*message.User = make(map[int]*message.User, 10)

var CurUser model.CurUser

func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id, _ := range OnlineUser {
		fmt.Println("用户id:\t", id)
	}
}

// 编写一个方法处理返回的notifyUserStatusMes的信息
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := OnlineUser[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	OnlineUser[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
