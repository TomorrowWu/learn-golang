package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//这里循环接收客户端发送的数据
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		//1,等待客户端通过conn发送信息
		//2,如果客户端没有write[发送],那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Println("服务器的read err=", err)
			return
		}
		//3,显示客户端发送的内容到服务器的终端
		content := string(buf[:n])
		fmt.Printf("客户端:%s\n", content)

		dict := map[string]string{
			"who are you": "我是小冰",
			"你的性别?":       "你猜猜看",
			"你会什么":        "我会讲故事",
			"你讲个吧":        "从前有座山",
		}
		res, ok := dict[content]
		if !ok {
			res = "你说啥"
		}
		conn.Write([]byte(res))
	}
}

func main() {
	fmt.Println("服务器开始监听")

	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listener.Close()

	for {
		fmt.Println("等待客户端来链接....")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
			continue
		} else {
			fmt.Printf("Accept() suc conn=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备一个协程,为客户端服务
		go process(conn)
	}

}
