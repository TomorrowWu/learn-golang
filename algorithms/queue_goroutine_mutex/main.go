package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

//1. 	1)创建一个数组模拟队列，每隔一定时间[随机的]， 给该数组添加一个数。
//2. 	2)启动两个协程，每隔一定时间(时间随机)到队列取出数据
//3. 	3)在控制台输出
//x号协程 服务 ---》x号客户
//x号协程 服务 ---》x号客户
//x号协程 服务 ---》x号客户
//4. 	4)使用锁机制即可。

var mutex sync.Mutex

type CircleQueue struct {
	maxSize int //4
	array   [10]int
	head    int //指向队列队首 0,包含元素
	tail    int //指向队尾 0,不含最后元素
}

func (circleQueue *CircleQueue) Push(val int) (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	if circleQueue.IsFull() {
		return errors.New("queue is full")
	}
	//queue.tail在队列尾部,不包含最后的元素
	circleQueue.array[circleQueue.tail] = val //把值给尾部
	circleQueue.tail = (circleQueue.tail + 1) % circleQueue.maxSize
	return
}

func (circleQueue *CircleQueue) Pop() (val int, err error) {
	mutex.Lock()
	defer mutex.Unlock()
	if circleQueue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	val = circleQueue.array[circleQueue.head]
	circleQueue.head = (circleQueue.head + 1) % circleQueue.maxSize

	return val, nil
}

func (circleQueue *CircleQueue) ListQueue() {
	mutex.Lock()
	defer mutex.Unlock()
	//取出当前队列有多少个元素
	size := circleQueue.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}

	//设计一个辅助的变量,指向head
	tempHead := circleQueue.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, circleQueue.array[tempHead])
		tempHead = (tempHead + 1) % circleQueue.maxSize
	}
	fmt.Println()
}

func (circleQueue *CircleQueue) IsFull() bool {
	//会预留一个空位,
	return (circleQueue.tail+1)%circleQueue.maxSize == circleQueue.head
}

func (circleQueue *CircleQueue) IsEmpty() bool {
	return circleQueue.tail == circleQueue.head
}

//取出环形队列有多少个元素
func (circleQueue *CircleQueue) Size() int {
	//这是一个关键的算法.
	//演算一下,size不可能超过maxSize,存在tail<head的情况,所以直接+maxSize,再取模
	return (circleQueue.tail + circleQueue.maxSize - circleQueue.head) % circleQueue.maxSize
}

func process(n int, queue *CircleQueue) {
	for {
		time.Sleep(3 * time.Second)
		val, err := queue.Pop()
		if err != nil {
			fmt.Println("err=", err)
		} else {
			fmt.Printf("%d号协程 服务-->> %d号客户\n", n, val)
		}
	}
}

func main() {

	//初始化一个环形队列
	queue := &CircleQueue{
		maxSize: 10,
		head:    0,
		tail:    0,
	}

	go process(1, queue)
	go process(2, queue)

	//每隔一定时间(随机),给数组添加一个数
	i := 0
	for {
		err := queue.Push(i)
		if err != nil {
			fmt.Println("err=", err)
		} else {
			i++
		}
		//interval := rand.Intn(3)
		time.Sleep(1 * time.Second)
	}
}
