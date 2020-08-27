package main

import (
	"fmt"
	"chat/client/handle"

)

var (
	userid int
	pwd string
)

func Verr(err error) {
	if err != nil {
		fmt.Println("发现错误：", err)
	}
}

func inputInfo() {
	fmt.Println("请输入id")
	fmt.Scanf("%d\n", &userid)
	fmt.Println("请输入密码")
	fmt.Scanf("%s\n", &pwd)

	uh := &handle.UserHandle{}


	err := uh.Login(userid, pwd)
	Verr(err)
}


func SelectMenu(key int) bool {
	var loop bool = true
	switch key {
		case 1:
			inputInfo()
			loop = false
			return loop
		case 2:
			fmt.Println("---正在注册新账号---")
			fmt.Println("请输入id：")
			fmt.Scanf("%d\n", &userid)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &pwd)


			uh := &handle.UserHandle{}
			err := uh.Register(userid, pwd)
			Verr(err)


			loop = false
			return loop
		case 3:
			fmt.Println("333")
			loop = false
			return loop
		default:
			fmt.Println("输入错误")
			return loop
	}
}

func main() {
	var key int
	for {
		fmt.Println("Ourchat v1.0")
		fmt.Println("1、登录")
		fmt.Println("2、注册")
		fmt.Println("3、退出")
		fmt.Println("输入数字，选择功能（1-3）")
		fmt.Scanf("%d\n", &key)
		loop := SelectMenu(key)
		if loop {
			continue
		} else {
			break
		}
	}
	fmt.Println("结束")
}
