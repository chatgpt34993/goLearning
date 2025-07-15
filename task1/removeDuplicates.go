package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    k := 1 // 初始时第一个元素总是唯一的
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[k-1] {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}

func main() {
    testCases := [][]int{
        {1, 1, 2},
        {0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
        {1, 2, 3, 4, 5},
        {},
        {1},
        {1, 1, 1, 1},
    }
    for _, tc := range testCases {
        // 复制原数组以便输出原始数据
        original := make([]int, len(tc))
        copy(original, tc)
        k := removeDuplicates(tc)
        fmt.Printf("输入: %v\n处理后前k个元素: %v\nk = %d\n\n", original, tc[:k], k)
    }
}