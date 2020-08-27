package authpkg

import (
	"errors"
)

var (
	ERROR_USER_PWD = errors.New("密码错误")
	ERROR_USER_NOEXIST = errors.New("用户不存在")
	ERROR_USER_EXIST = errors.New("用户已存在")
)
