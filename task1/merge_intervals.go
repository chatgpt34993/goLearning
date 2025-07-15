package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for _, current := range intervals[1:] {
		last := result[len(result)-1]
		if current[0] <= last[1] {
			// 重叠则合并
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 不重叠则添加新区间
			result = append(result, current)
		}
	}
	return result
}

func main() {
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {2, 3}},
		{{1, 2}, {3, 4}, {5, 6}},
		{},
	}
	for _, tc := range testCases {
		fmt.Printf("输入: %v\n输出: %v\n\n", tc, merge(tc))
	}
}    