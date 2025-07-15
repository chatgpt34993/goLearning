package main

import "fmt"

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if idx, exists := numMap[complement]; exists {
            return []int{idx, i}
        }
        numMap[num] = i
    }
    return nil
}

func main() {
    testCases := []struct {
        nums   []int
        target int
    }{
        {[]int{2, 7, 11, 15}, 9},
        {[]int{3, 2, 4}, 6},
        {[]int{3, 3}, 6},
        {[]int{-1, -2, -3, -4}, -7},
    }

    for _, tc := range testCases {
        result := twoSum(tc.nums, tc.target)
        fmt.Printf("输入: nums=%v, target=%d\n输出: %v\n\n", tc.nums, tc.target, result)
    }
}