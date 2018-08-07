package singlequeue

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int //数组是一种数据类型,数据结构是对数据+算法
	front   int    //表示指向队列首
	rear    int    //表示指向队列的尾部
}

//添加数据到队列
func (queue *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if queue.rear == queue.maxSize-1 { //rear是队列尾部(含最后元素)
		return errors.New("queue full")
	}

	queue.rear++
	queue.array[queue.rear] = val
	return
}

func (queue *Queue) GetQueue() (val int, err error) {

}

//显示队列
func (queue *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是:")
	//queue.front不包含队首元素
	for i := queue.front + 1; i < queue.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, queue.array[i])
	}
}

func main() {
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		array:   [5]int{},
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("3. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列表的数")
			fmt.Scanln(&val)

			err := queue.AddQueue(val)
			if err == nil {
				fmt.Println("加入队列OK")
			} else {
				fmt.Println(err.Error())
			}
		case "get":
			fmt.Println("get")
			val, err := queue.GetQueue()
			if err == nil {

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
