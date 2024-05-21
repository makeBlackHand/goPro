package model

import (
	"goPro/BasicSyntax/chatRoom/common/message"
	"net"
)

type CurUser struct {
	Con net.Conn
	message.User
}
