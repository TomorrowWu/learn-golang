package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

//Transfer 这里将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //传输时,使用缓冲
}

func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {
	_, err = transfer.Conn.Read(transfer.Buf[:4]) //从conn读取
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[:4])

	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}

	//pkgLen 反序列化 -> message.Message
	err = json.Unmarshal(transfer.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte

	//binary包实现了简单的数字与字节序列的转换以及变长值的编解码
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	//发送长度
	n, err := transfer.Conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail err=", err)
		return
	}

	_, err = transfer.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail err=", err)
		return
	}
	return
}
