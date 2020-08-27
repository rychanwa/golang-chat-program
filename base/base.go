package base

import (
	// "fmt"
)

const (

	LoginType = "login"
	LoginresType = "loginres"

	RegisterType = "reg"
	RegisterresType = "regres"

)

// 消息结构体
type Message struct {
	MessType string `json:"messtype"`
	MessData string `json:"messdata"`
}

// 用户登录数据
type UserArgs struct {
	UserId int `json:"userid"`
	Upwd string `json:"upwd"`
}

type LoginRes struct {
	Code int `json:"code"` // 登录状态码
	Error string `json:"error"` // 返回错误信息
	Uids []int `json:"uids"` // 在线用户
}

type RegisterUser struct {
	User UserInfo `json:"user"`
}

type RegRes struct {
	Code int `json:"code"` // 状态码
	Error string `json:"error"` // 返回错误信息
}
