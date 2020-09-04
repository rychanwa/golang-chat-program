package main

import(
	"fmt"
	"net"
	"chat/base"
	"chat/server/cprocess"
	"chat/server/utils"
	"io"

)

type Processor struct {
	Con net.Conn
}
// 按照类型处理消息
func (this *Processor) OperaMessage(mess base.Message) (err error) {
	// fmt.Println("测试消息内容：", mess)

	switch mess.MessType {
		case base.LoginType:

			up := &cprocess.Userprocess{
				Con : this.Con,
			}
			up.CheckLogin(mess)

		case base.RegisterType:

			up := &cprocess.Userprocess{
				Con : this.Con,
			}
			up.RegistUser(mess)

		case base.SmsType:
			// 群发消息
			up := &cprocess.GroupSms{}
			up.SendGroupMes(mess)



		default:
			fmt.Println("未知消息类型")
	}
	return
}

func (this *Processor) MainProcess() (err error) {
	for {
		tf := &utils.Transfer{
			Con : this.Con,
		}
		mess, err := tf.ReadMessage()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端主动断开了连接。")
			} else {
				fmt.Println("错误：", err)

			}
			return err
		}
		fmt.Println(mess)
		err = this.OperaMessage(mess)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}
