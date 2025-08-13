package main

import (
	"fmt"
	"time"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。

题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

func printOdds() {
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
	}
}

func printEvens() {
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
	}
}

func taskSchedule(functions []func()) {
	done := make(chan time.Duration)

	for _, function := range functions {
		go func(f func()) {
			start := time.Now()
			f()
			done <- time.Since(start)
		}(function)
	}

	// 只接收与任务数量相同的结果
	for i := 0; i < len(functions); i++ {
		value := <-done
		fmt.Printf("执行耗时: %v\n", value)
	}
	defer close(done)
}

func task(num int) {
	time.Sleep(time.Duration(num) * time.Second)
}

func main() {
	//go printOdds()
	//go printEvens()
	//
	//// 等待足够长的时间让两个协程完成
	//time.Sleep(1 * time.Second)
	//fmt.Println("主协程结束")

	fmt.Println("====================")
	start := time.Now()
	taskSlice := []func(){
		func() { task(1) },
		func() { task(2) },
	}
	taskSchedule(taskSlice)
	fmt.Println(time.Since(start))
}
