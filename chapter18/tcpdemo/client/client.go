package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	defer conn.Close()
	//fmt.Println("conn 成功=", conn)

	reader := bufio.NewReader(os.Stdin) //标准输入(终端)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		//退出
		line = strings.Trim(line, "\r\n") //去掉空格
		if line == "exit" {
			return
		}

		if len(line) == 0 {
			continue
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("client write err=", err)
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Println("服务器的read err=", err)
			return
		}
		content := string(buf[:n])
		fmt.Println("服务器回复:", content)
	}

}
