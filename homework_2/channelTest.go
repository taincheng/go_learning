package main

import (
	"fmt"
	"time"
)

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/

func test1() {
	numCh := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("send value: %d\n", i)
			numCh <- i
		}
	}()

	go func() {
		for value := range numCh {
			fmt.Printf("get value: %d\n", value)
		}
	}()
}

func test2() {
	numCh := make(chan int, 5)

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Printf("send value: %d\n", i)
			numCh <- i
		}
	}()

	go func() {
		for value := range numCh {
			fmt.Printf("get value: %d\n", value)
		}
	}()
}

func main() {
	test1()
	test2()
	time.Sleep(1 * time.Second)
}
