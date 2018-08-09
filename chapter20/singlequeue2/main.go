package main

import (
	"errors"
	"fmt"
	"os"
)

//创建一个链表模拟队列，实现 数据入队列，数据出队列，显示队列.

type Node struct {
	no   int
	next *Node //这个表示指向下一个结点
}

type Queue struct {
	head *Node //表示头结点
}

func (queue *Queue) AddQueue(newNode *Node) {
	temp := queue.head
	if temp == nil {
		queue.head = newNode
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
}

func (queue *Queue) GetQueue() (node *Node, err error) {
	if queue.head == nil {
		return nil, errors.New("queue empty")
	}

	temp := queue.head

	queue.head = temp.next

	temp.next = nil
	return temp, nil
}

func (queue *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是:")
	temp := queue.head
	if temp == nil {
		return
	}
	for {
		fmt.Printf("%d\t", temp.no)
		temp = temp.next
		if temp == nil {
			break
		}
	}
	fmt.Println()
}

func main() {
	queue := &Queue{
		head: nil,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列表的数")
			fmt.Scanln(&val)

			queue.AddQueue(&Node{
				no:   val,
				next: nil,
			})
		case "get":
			val, err := queue.GetQueue()
			if err == nil {
				fmt.Println("从队列中取出了一个数=", val)
			} else {
				fmt.Println(err.Error())
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
