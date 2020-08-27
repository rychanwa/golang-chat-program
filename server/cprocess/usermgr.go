package cprocess

import (
	"fmt"
)

var userlist *UserList // 定义了结构体的指针的变量

type UserList struct {
	Users map[int]*Userprocess
}

func init() {
	userlist = &UserList{
		Users : make(map[int]*Userprocess, 1024),
	}
}

func (this *UserList) AddUserList(up *Userprocess) {
	this.Users[up.Uid] = up
}

func (this *UserList) DelUserList(userid int) {
	delete(this.Users, userid)
}

func (this *UserList) GetAllUsers() map[int]*Userprocess {
	return this.Users
}

func (this *UserList) GetUserById(userid int) (up *Userprocess, err error) {
	up, ok := this.Users[userid]
	if !ok {
		err = fmt.Errorf("用户%d，当前不在线", userid)
		return
	}
	return
}
