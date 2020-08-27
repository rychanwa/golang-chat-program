package handle

import (
	"fmt"
	"os"
	"net"
	"chat/client/utils"
)

func ShowMenu() {
	fmt.Println("--- 登录成功 ---")
	fmt.Println(">> 1 查看在线好友 <<")
	fmt.Println(">> 2 发送消息 <<")
	fmt.Println(">> 3 消息列表 <<")
	fmt.Println(">> 4 退出系统 <<")
	fmt.Println("输入数字(1-4)选择功能")

	var key int
	fmt.Scanf("%d\n", &key)

	switch key {
		case 1:
			fmt.Println("")
		case 2:
			fmt.Println("")
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

		fmt.Println("通讯内容：", mes)
	}
}
