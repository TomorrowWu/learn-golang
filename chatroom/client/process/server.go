package process

import (
	"fmt"
	"os"
)

func ShowMenu() {
	fmt.Println("-------------恭喜xxx登录成功-------------")
	fmt.Println("-------------1 显示在线用户列表-------------")
	fmt.Println("-------------2 发送消息-------------")
	fmt.Println("-------------3 信息列表-------------")
	fmt.Println("-------------4 退出系统-------------")
	fmt.Println("请选择(1-4):")

	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统")
		//可以先通知服务器,做一些操作,入库操作
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确")
	}
}
