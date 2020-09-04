package base

import ()

type UserInfo struct {
	UserId int `json:"userid"`
	Upwd string `json:"upwd"`
	UserStatus int `json:"userstatus"` // 用户的在线状态
}
