package main

import (
	"fmt"
	"sort"
)

/*
56. 合并区间 [https://leetcode.cn/problems/merge-intervals/description/]
*/
func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	index := 0
	for i := 1; i < len(intervals); i++ {
		if !(res[index][1] < intervals[i][0] || res[index][0] > intervals[i][1]) {
			res[index][0] = min(res[index][0], intervals[i][0])
			res[index][1] = max(res[index][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
			index++
		}
	}
	return res
}

func main() {
	args := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	fmt.Println(merge(args))
}
