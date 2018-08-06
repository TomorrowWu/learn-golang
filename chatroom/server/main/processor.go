package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	process2 "go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (processor *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		up := &process2.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessLogin(mes)
		if err != nil {
			return
		}
	case message.RegisterMesType:
		//处理注册
		//处理注册
		up := &process2.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessRegister(mes) // type : data
	case message.SmsMesType:
		//创建一个SmsProcess实例完成转发群聊消息.
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在,无法处理...")
	}
	return
}

func (processor *Processor) process2() (err error) {
	tf := &utils.Transfer{
		Conn: processor.Conn,
		Buf:  [8096]byte{},
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务端也退出..")
			} else {
				fmt.Println("read pkg err=", err)
			}
			//客户端出现错误,就断开和客户端的链接
			return err
		}

		err = processor.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
