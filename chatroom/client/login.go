package main

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

//写一个函数,完成登录
func login(userId int, userPwd string) (err error) {
	//下一步开始定协议...
	//fmt.Printf("userId = %d userPwd=%s", userId, userPwd)

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//2,conn发送消息
	var mes message.Message
	mes.Type = message.LoginMesType

	//3,创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4,序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5,mes.Data字段
	mes.Data = string(data)

	//6,mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	err = writePkg(conn, data)
	if err != nil {
		return
	}

	//处理服务器登录返回的消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}
