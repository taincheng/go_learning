package main

import (
	"fmt"
	"unsafe"
)

/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/

func add10(num *int) {
	*num += 10
}

func slicePointerMultiplyOf2(intSlice *[]int) {
	for i := range *intSlice {
		(*intSlice)[i] *= 2
	}
	// 修改原数组切片的结构，地址不变
	*intSlice = append(*intSlice, 2)
}

func sliceMultiplyOf2(intSlice []int) {
	for i := range intSlice {
		intSlice[i] *= 2
	}
	// 修改原数组切片的结构，地址不变
	intSlice = append(intSlice, 2)
}

func main() {
	num := 10
	add10(&num)
	fmt.Println(num)

	intSlice := []int{1}
	// intSlice切片的地址不会发生变化，因为切片是结构体，只要结构体的原变量不发生变化，地址就固定
	fmt.Printf("Before: %p, Data: %p\n", &intSlice, unsafe.Pointer(&(intSlice)[0]))
	slicePointerMultiplyOf2(&intSlice)
	fmt.Println(intSlice)
	// append 一个元素，触发了扩容，切片中的data需要扩大，所以换了新的内存地址
	fmt.Printf("Before: %p, Data: %p\n", &intSlice, unsafe.Pointer(&(intSlice)[0]))

	// 传递切片，只是值传递，传递了整个slice结构体，因此函数体内的slice append操作，不会影响外部的slice
	// intSlice 依然是 *2 的效果，是因为slice是共享同一个底层数组，底层使用了指针取值进行 *2 操作
	sliceMultiplyOf2(intSlice)
	fmt.Println(intSlice)
	fmt.Printf("Before: %p, Data: %p\n", &intSlice, unsafe.Pointer(&(intSlice)[0]))
}
