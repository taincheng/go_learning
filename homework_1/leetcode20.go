package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

1.左括号必须用相同类型的右括号闭合。
2.左括号必须以正确的顺序闭合。
3.每个右括号都有一个对应的相同类型的左括号。
*/
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	strMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []byte
	for _, str := range s {
		if len(stack) == 0 {
			stack = append(stack, byte(str))
		} else {
			if strMap[byte(str)] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, byte(str))
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(])["))
}
