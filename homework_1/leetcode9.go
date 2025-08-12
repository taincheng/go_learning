package main

import (
	"fmt"
	"strconv"
)

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
*/
func isPalindrome(x int) bool {
	byteValue := []byte(strconv.Itoa(x))
	length := len(byteValue)
	i := 0
	j := length - 1
	for i <= j {
		if byteValue[i] != byteValue[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func isPalindrome1(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNum := 0
	for x > revertedNum {
		revertedNum = revertedNum*10 + x%10
		x /= 10
	}
	return x == revertedNum || x == revertedNum/10
}
func main() {
	fmt.Println(isPalindrome1(121))
}
