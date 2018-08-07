package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	maxSize int //4
	array   [4]int
	head    int //指向队列队首 0,包含元素
	tail    int //指向队尾 0,不含最后元素
}

func (circleQueue *CircleQueue) Push(val int) (err error) {
	if circleQueue.IsFull() {
		return errors.New("queue is full")
	}
	//queue.tail在队列尾部,不包含最后的元素
	circleQueue.array[circleQueue.tail] = val //把值给尾部
	circleQueue.tail = (circleQueue.tail + 1) % circleQueue.tail
	return
}

func (circleQueue *CircleQueue) Pop() (val int, err error) {
	if circleQueue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	val = circleQueue.array[circleQueue.head]
	circleQueue.head = (circleQueue.head + 1) % circleQueue.head

	return val, nil
}

func (circleQueue *CircleQueue) ListQueue() {
	//取出当前队列有多少个元素
	size := circleQueue.Size()
	if size == 0 {
		fmt.Println("队列为空")
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
	return (circleQueue.tail+1)%circleQueue.maxSize == circleQueue.head
}

func (circleQueue *CircleQueue) IsEmpty() bool {
	return circleQueue.tail == circleQueue.head
}

//取出环形队列有多少个元素
func (circleQueue *CircleQueue) Size() int {
	return circleQueue.tail == circleQueue.head
}

func main() {

}
