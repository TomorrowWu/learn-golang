package main

import (
	"fmt"
	"go_code/chatroom/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {
	//这里循环接收客户端发送的数据
	defer conn.Close()

	//创建一个总控,调用
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯协程错误=", err)
		return
	}
}

func initUserDao() {
	//这里需要注意一个初始化顺序的问题,
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()

	//
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端来连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			continue
		}

		//一旦连接成功,则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
