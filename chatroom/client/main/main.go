package main

import (
	"fmt"
	"go_code/chatroom/client/process"
	"os"
)

var (
	userId  int
	userPwd string
)

func main() {
	//接收用户的选择
	var key int

	for true {
		fmt.Println("---------------------欢迎登陆多人聊天系统--------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId) //需要有回车符(\n)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)

			userProcess := &process.UserProcess{}
			userProcess.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误,请重新输入")
		}
	}
}
