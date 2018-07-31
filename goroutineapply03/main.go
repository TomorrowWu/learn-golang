package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}

	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//var num int
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//不是素数
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum 协程取不到数据,退出")
	//其他协程可能还在处理,不能关闭primeChan

	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //放结果

	exitChan := make(chan bool, 4)

	start := time.Now().UnixNano()

	go putNum(intChan)

	for i := 1; i <= 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {

		for i := 1; i <= 4; i++ {
			<-exitChan
		}

		end := time.Now().UnixNano()

		fmt.Println("总耗时=", end-start)

		close(primeChan)
	}()

	for {
		//res, ok := <-primeChan
		_, ok := <-primeChan
		if !ok {
			break
		}

		//fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main退出")
}
