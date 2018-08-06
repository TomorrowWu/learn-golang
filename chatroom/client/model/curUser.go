package model

import (
	"go_code/chatroom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
