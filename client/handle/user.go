package handle

import (
	"fmt"
	"encoding/json"
	"net"
	"chat/base"
	"encoding/binary"
	"chat/client/utils"
)

type UserHandle struct {}


func (this *UserHandle) Register(userid int, pwd string) (err error) {
	con, err := net.Dial("tcp", "localhost:8856")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var mess base.Message
	mess.MessType = base.RegisterType

	var regist base.RegisterUser
	regist.User.UserId = userid // 结构体继承，字段会自动追寻
	regist.User.Upwd = pwd

	data, err := json.Marshal(regist)
	if err != nil {
		fmt.Println("序列化注册信息：", err)
		return err
	}

	mess.MessData = string(data)
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tf := &utils.Transfer{
		Con : con,
	}
	err = tf.SendMessage(data)
	if err != nil {
		fmt.Println("发送注册信息出错：", err)
	}

	mes, err := tf.ReadMessage()
	if err != nil {
		fmt.Println("读取注册返回信息出错，", err)
	}
	fmt.Println("注册结果：", mes)

	var regres base.RegRes
	err = json.Unmarshal([]byte(mes.MessData), &regres)


	if err != nil {
		fmt.Println("反序列化注册结果时发生错误：", err)
	} else {
		fmt.Println("注册成功", regres)
	}

	return

}

func (this *UserHandle) Login(userid int, pwd string) (err error) {

	con, err := net.Dial("tcp", "localhost:8856")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var mess base.Message
	mess.MessType = base.LoginType

	var userargs base.UserArgs
	userargs.UserId = userid
	userargs.Upwd = pwd

	data, err := json.Marshal(userargs)
	if err != nil {
		fmt.Println(err)
		return err
	}

	mess.MessData = string(data)
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var messLen uint32
	messByte := make([]byte, 4)
	messLen = uint32(len(data))
	binary.BigEndian.PutUint32(messByte[:], messLen)
	n, err := con.Write(messByte) // 发送消息长度
	if err != nil || n != 4 {
		fmt.Println(err)
		return err
	}

	// fmt.Println("客户端发送的消息内容是：", string(data))

	// 发送消息
	_, err = con.Write(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 开始接收登录反馈的信息
	tf := &utils.Transfer{
		Con : con,
	}
	mes, err := tf.ReadMessage()

	if err != nil {
		fmt.Println("读取登陆返回信息出错，", err)
	}
	fmt.Println("登录结果：", mes)

	var loginres base.LoginRes
	err = json.Unmarshal([]byte(mes.MessData), &loginres)
	if err != nil {
		fmt.Println("反序列化登录反馈信息出错", err)
	} else {
		// fmt.Println(loginres)
		if loginres.Code == 900 {

			// 显示在线用户
			for _, value := range loginres.Uids {

				if value == userid {
					// 从列表中排除自己
					continue
				}
				fmt.Println(value)

			}
			fmt.Println("\n\n")



			go MessHandle(con)
			for {
				ShowMenu()
			}
		}
	}

	defer con.Close()
	return nil
}
