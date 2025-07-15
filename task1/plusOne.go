package main

import "fmt"

func plusOne(digits []int) []int {
    n := len(digits)
    for i := n - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    // 如果所有位都是9，需要进位
    result := make([]int, n+1)
    result[0] = 1
    return result
}

func main() {
    testCases := [][]int{
        {1, 2, 3},
        {4, 3, 2, 1},
        {9},
        {9, 9, 9},
        {5, 6, 7, 8, 9},
        {0},
    }
    for _, tc := range testCases {
        fmt.Printf("输入: %v\n输出: %v\n\n", tc, plusOne(tc))
    }
}