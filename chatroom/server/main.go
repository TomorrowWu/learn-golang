package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4]) //从conn读取
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}

	//pkgLen 反序列化 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}
func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte

	//binary包实现了简单的数字与字节序列的转换以及变长值的编解码
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail err=", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail err=", err)
		return
	}
	return
}

func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		err = serverProcessLogin(conn, mes)
		if err != nil {
			return
		}
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在,无法处理...")
	}
	return
}
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1,声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2,声明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500 //500 状态码表示该用户不存在
		loginResMes.Error = "该用户不存在,请注册再使用..."
	}

	//3,将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	err = writePkg(conn, data)
	if err != nil {
		return
	}

	return
}

func process(conn net.Conn) {
	//这里循环接收客户端发送的数据
	defer conn.Close()
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务端也退出..")
			} else {
				fmt.Println("read pkg err=", err)
			}
			//客户端出现错误,就断开和客户端的链接
			return
		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

func main() {

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
