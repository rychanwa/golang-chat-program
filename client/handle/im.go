package handle

import (
	"fmt"
	"chat/base"
	"encoding/json"
	"chat/client/utils"
)

type SmsHandle struct {}

func (this *SmsHandle) SendMessForGroup(content string) (err error) {
	var messbody base.Message
	messbody.MessType = base.SmsType

	var messcon base.Sms
	messcon.SmsCon = content
	messcon.UserId = curuser.UserId
	messcon.UserStatus = curuser.UserStatus

	data, err := json.Marshal(messcon)
	if err != nil {
		fmt.Println("群发消息序列化失败，错误：", err.Error())
		return
	}

	messbody.MessData = string(data)

	data, err = json.Marshal(messbody)
	if err != nil {
		fmt.Println("群发消息主体序列化失败，错误：", err.Error())
		return
	}

	tf := &utils.Transfer{
		Con : curuser.Con,
	}

	err = tf.SendMessage(data)
	if err != nil {
		fmt.Println("群消息发送失败，错误：", err.Error())
		return
	}
	return
}
