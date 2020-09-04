package cprocess

import (
	"fmt"
	"chat/base"
	"net"
	"chat/server/utils"
	"encoding/json"
)

type GroupSms struct {}

func (this *GroupSms) SendGroupMes(mes base.Message) {

	var sms base.Sms
	err := json.Unmarshal([]byte(mes.MessData), &sms)
	if err != nil {
		fmt.Println("反序列化消息内容失败，错误：", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化消息失败，错误：", err)
		return
	}

	for id, up := range userlist.Users {
		if id == sms.UserId {
			continue
		}
		//
		this.SendMessageForUser(data, up.Con)
	}
}

func (this *GroupSms) SendMessageForUser(data []byte, con net.Conn) {
	//
	tf := &utils.Transfer{
		Con : con,
	}

	err := tf.SendMessage(data)
	if err != nil {
		fmt.Println("群发消息失败，错误：", err)

	}
}
