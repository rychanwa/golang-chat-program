package model

import (
	"net"
	"chat/base"
)

type CurUser struct {
	Con net.Conn
	base.UserInfo
}
