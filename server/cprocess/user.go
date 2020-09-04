package cprocess

import (
	"net"
	"chat/base"
	"chat/server/utils"
	"encoding/json"
	"fmt"
	"chat/server/model"
)

type Userprocess struct {
	Con net.Conn
	Uid int
}


// 通知用户在线
func (this *Userprocess) NotifyOtherUserOnline(userid int) {
	for id, up := range userlist.Users {
		if id == userid {
			continue
		}
		up.OnlineNotify(userid)
	}

}

func (this *Userprocess) OnlineNotify(userid int) {
	//
	var mes base.Message
	mes.MessType = base.OnlineNotifyType

	var notify base.UserOnlineNotify
	notify.Userid = userid
	notify.Userstatus = base.UserOnline

	data, err := json.Marshal(notify)
	if err != nil {
		fmt.Println("通知序列化失败，错误：", err)
		return
	}

	mes.MessData = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("通知消息序列化失败，错误：", err)
		return
	}

	tf := &utils.Transfer{
		Con : this.Con,

	}

	err = tf.SendMessage(data)
	if err != nil {
		fmt.Println("在线通知发送失败，错误：", err)
		return
	}


}

// 注册用户
func (this *Userprocess) RegistUser(mess base.Message) (err error) {
	//
	var registerInfo base.RegisterUser
	err = json.Unmarshal([]byte(mess.MessData), &registerInfo)
	if err != nil {
		fmt.Println("反序列化注册用户时，出错：", err)
		return
	}
	//
	var mainMess base.Message
	mainMess.MessType = base.RegisterresType
	var conMess base.RegRes


	err = authpkg.ConstDao.Register(&registerInfo.User)
	if err != nil {
		if err == authpkg.ERROR_USER_EXIST {
			conMess.Code = 903
			conMess.Error = err.Error()
		} else {
			conMess.Code = 903
			conMess.Error = "注册失败"
		}


	} else {
		conMess.Code = 902
		conMess.Error = "注册成功"

	}

	data, err := json.Marshal(conMess)
	if err != nil {
		fmt.Println("序列化返回注册内容信息时，失败：", err)
		return
	}
	mainMess.MessData = string(data)
	data, err = json.Marshal(mainMess)
	if err != nil {
		fmt.Println("序列化返回注册信息时，失败：", err)
		return
	}
	// 返回登录后的消息
	transfer := &utils.Transfer{
		Con : this.Con,

	}
	err = transfer.SendMessage(data)
	return
}

// 验证登录
func (this *Userprocess) CheckLogin(mess base.Message) (err error) {
	var loginInfo base.UserArgs
	err = json.Unmarshal([]byte(mess.MessData), &loginInfo)
	if err != nil {
		fmt.Println("反序列化登录信息时，出错：", err)
		return
	}
	var mainMess base.Message
	mainMess.MessType = base.LoginresType
	var conMess base.LoginRes

	user, err := authpkg.ConstDao.CheckUserInfo(loginInfo.UserId, loginInfo.Upwd)

	if err != nil {

		if err == authpkg.ERROR_USER_PWD {
			conMess.Code = 901
			conMess.Error = err.Error()
		} else if err == authpkg.ERROR_USER_NOEXIST {
			conMess.Code = 901
			conMess.Error = err.Error()
		} else {
			conMess.Code = 901
			conMess.Error = "未知错误"
		}


	} else {
		conMess.Code = 900
		conMess.Error = "login success"

		// 写入在线用户id
		this.Uid = loginInfo.UserId

		userlist.AddUserList(this) // 注意这里，在usermgr里面定义了变量userlist，类型是结构体UserList的指针

		for key, _ := range userlist.Users {
			conMess.Uids = append(conMess.Uids, key)


		}
		this.NotifyOtherUserOnline(loginInfo.UserId)
		fmt.Println("当前登录用户信息：", user)
	}


	/*if loginInfo.Upwd == "chen" && loginInfo.UserId == 123456 {
		conMess.Code = 900
		conMess.Error = "login success"
	} else {
		conMess.Code = 901
		conMess.Error = "login fail"
	}*/
	data, err := json.Marshal(conMess)
	if err != nil {
		fmt.Println("序列化验证信息时，失败：", err)
		return
	}
	mainMess.MessData = string(data)
	data, err = json.Marshal(mainMess)
	if err != nil {
		fmt.Println("序列化返回信息时，失败：", err)
		return
	}
	// 返回登录后的消息
	transfer := &utils.Transfer{
		Con : this.Con,

	}
	err = transfer.SendMessage(data)
	return
}
