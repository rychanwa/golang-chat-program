package handle

import (
	"fmt"
	"chat/base"
	"chat/client/model"
)

var OnlineUser map[int]*base.UserInfo = make(map[int]*base.UserInfo, 11)
var curuser model.CurUser

func ViewOnlineUserList() {
	fmt.Println("当前在线用户")
	for id, _ := range OnlineUser {
		fmt.Println("在线用户id：", id)
	}
}

func UpdateUserStatus(uon *base.UserOnlineNotify) {
	user, ok := OnlineUser[uon.Userid]
	if !ok {
		user = &base.UserInfo{
			UserId : uon.Userid,
		}
	}
	user.UserStatus = uon.Userstatus
	OnlineUser[uon.Userid] = user
	ViewOnlineUserList()
}
