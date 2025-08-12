package main

import "fmt"

/*
26. 删除有序数组中的重复项
*/
func removeDuplicates(nums []int) int {
	length := len(nums)
	i := 1
	for j := 1; j < length; j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	args := []int{0, 1, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(args))
}
