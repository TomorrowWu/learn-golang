package main

import "fmt"

//解决Josephu问题

type CatNode struct {
	no   int //猫猫的编号
	next *CatNode
}

type Link struct {
	head *CatNode
}

func (link *Link) InsertCatNode(newCatNode *CatNode) {

	//判断是不是添加第一只猫
	if link.head == nil {
		link.head = newCatNode
		newCatNode.next = newCatNode //构成一个环形
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	//定义一个临时变量，帮忙,找到环形的最后结点
	temp := link.head
	for {
		if temp.next == link.head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = link.head
}

//输出这个环形的链表
func (link *Link) ListCircleLink() {
	fmt.Println("环形链表的情况如下：")
	temp := link.head
	if temp == nil {
		fmt.Println("空空如也的环形链表...")
		return
	}
	for {
		fmt.Printf("猫的信息为=[no=%d] ->\n", temp.no)
		if temp.next == link.head {
			break
		}
		temp = temp.next
	}
}

func (link *Link) DelCatNode(id int) {
	temp := link.head
	helper := link.head
	//空链表
	if temp == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return
	}

	//如果只有一个结点
	if temp.next == link.head {
		if temp.no == id {
			//temp.next = nil
			link.head = nil
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起，没有no=%d\n", id)
		}
		return
	}

	//将helper 定位到链表最后
	for {
		if helper.next == link.head {
			break
		}
		helper = helper.next
	}

	//如果有两个包含两个以上结点
	flag := true
	//从head开始遍历
	for {
		if temp.next == link.head { //如果到这来，说明我比较到最后一个【最后一个还没比较】
			break
		}
		if temp.no == id {
			if temp == link.head {
				//说明删除的是头结点
				link.head = link.head.next //换head
			}
			helper.next = temp.next
			//temp.next = nil
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.next //移动 【比较】
		helper = helper.next
	}

	//这里还有比较一次
	if flag { //如果flag 为真，则我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
			//temp.next = nil
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起，没有no=%d\n", id)
		}
	}
}

func main() {

	link := &Link{}

	//创建一只猫
	cat1 := &CatNode{
		no: 1,
	}
	cat2 := &CatNode{
		no: 2,
	}
	cat3 := &CatNode{
		no: 3,
	}
	cat4 := &CatNode{
		no: 4,
	}
	cat5 := &CatNode{
		no: 5,
	}
	cat6 := &CatNode{
		no: 6,
	}
	cat7 := &CatNode{
		no: 7,
	}
	cat8 := &CatNode{
		no: 8,
	}

	link.InsertCatNode(cat1)
	link.InsertCatNode(cat2)
	link.InsertCatNode(cat3)
	link.InsertCatNode(cat4)
	link.InsertCatNode(cat5)
	link.InsertCatNode(cat6)
	link.InsertCatNode(cat7)
	link.InsertCatNode(cat8)
	link.ListCircleLink()

	//link.DelCatNode(30)
	//
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//link.ListCircleLink()
	//
	//link.DelCatNode(1)
	//
	//link.ListCircleLink()
	//
	//link.InsertCatNode(cat1)
	//
	//link.ListCircleLink()

	//n := 8 //人数
	k := 4
	m := 10
	var arr []int //出队编号序列

	//指针移到编号为4的元素
	temp := link.head
	for {
		if temp.no == k {
			break
		}
		temp = temp.next
	}

	for link.head != nil {
		for i := 1; i <= m; i++ {
			if i == m {
				//数到m,人出列
				//删除一个节点
				link.DelCatNode(temp.no)
				arr = append(arr, temp.no)
				break
			}
			temp = temp.next
		}
		temp = temp.next
	}

	fmt.Println("出队编号的序列arr=", arr)
}
