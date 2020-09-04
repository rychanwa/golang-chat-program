package handle

import (
	"fmt"
	"os"
	"net"
	"chat/client/utils"
	"chat/base"
	"encoding/json"
)

func ShowMenu() {
	fmt.Println("--- xxx登录成功 ---")
	fmt.Println(">> 1 查看在线好友 <<")
	fmt.Println(">> 2 发送消息 <<")
	fmt.Println(">> 3 消息列表 <<")
	fmt.Println(">> 4 退出系统 <<")
	fmt.Println("输入数字(1-4)选择功能")

	var key int
	fmt.Scanf("%d\n", &key)

	var content string

	sh := &SmsHandle{}

	switch key {
		case 1:
			ViewOnlineUserList()

		case 2:
			fmt.Println("请输入群发内容：")
			fmt.Scanf("%s\n", &content)
			sh.SendMessForGroup(content)
		case 3:
			fmt.Println("")
		case 4:
			fmt.Println("")
			os.Exit(1)
		default:
			fmt.Println("error")
	}
}

// 和服务器保持通讯
func MessHandle(con net.Conn) {
	tf := &utils.Transfer{
		Con : con,
	}

	for {
		fmt.Println("正在与服务器通讯中··· ···")
		mes, err := tf.ReadMessage()
		if err != nil {
			fmt.Println("server.go 通讯发生错误：", err)
			return
		}

		// fmt.Println("通讯内容：", mes)
		switch mes.MessType {
			case base.OnlineNotifyType:
				//
				var UserOnlineNotify base.UserOnlineNotify

				json.Unmarshal([]byte(mes.MessData), &UserOnlineNotify)
				UpdateUserStatus(&UserOnlineNotify)

			case base.SmsType:
				// 群发消息
				var sms base.Sms
				err = json.Unmarshal([]byte(mes.MessData), &sms)
				if err != nil {
					fmt.Println("server.go sms反序列化发生错误：", err)
					return
				}

				info := fmt.Sprintf("用户：\t %d， 对大家说：\t %s", sms.UserId, sms.SmsCon)
				fmt.Println(info)
				fmt.Println()



			default:
				fmt.Println("未知消息类型")

		}
	}
}
