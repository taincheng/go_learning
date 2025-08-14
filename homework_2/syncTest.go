package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

type counter interface {
	Add()
	Value() int64
}

type SafeCounter struct {
	mu      sync.Mutex
	counter int64
}

func (c *SafeCounter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

func (c *SafeCounter) Value() int64 {
	return c.counter
}

type AtomicCounter struct {
	counter int64
}

func (c *AtomicCounter) Add() {
	atomic.AddInt64(&c.counter, 1)
}

func (c *AtomicCounter) Value() int64 {
	return c.counter
}

func testAdd(counter counter) {
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				counter.Add()
			}
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	fmt.Println(counter.Value())
}

func main() {
	safeCounter := SafeCounter{}
	testAdd(&safeCounter)

	atomicCounter := AtomicCounter{}
	testAdd(&atomicCounter)
}
