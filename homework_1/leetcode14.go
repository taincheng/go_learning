package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""
*/
func longestCommonPrefix(strs []string) string {
	preStr := strs[0]
	for index := 1; index < len(strs); index++ {
		str := strs[index]
		length := min(len(str), len(preStr))
		if length == 0 {
			return ""
		} else {
			index := 0
			for index < length && preStr[index] == str[index] {
				index++
			}
			preStr = preStr[:index]
		}
	}
	return preStr
}
func main() {
	args := []string{
		"dog", "racecar", "car",
	}
	fmt.Println(longestCommonPrefix(args))
}
