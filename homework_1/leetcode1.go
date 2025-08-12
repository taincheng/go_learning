package main

import "fmt"

/*
1. 两数之和
*/
func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int, 8)
	for index, value := range nums {
		_, exist := mapNums[target-value]
		if exist {
			return []int{mapNums[target-value], index}
		} else {
			mapNums[value] = index
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
