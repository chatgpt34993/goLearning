package main

import "fmt"

// 先声明singleNumber函数
func singleNumber(nums []int) int {
    res := 0
    for _, num := range nums {
        res ^= num
    }
    return res
}

func main(){
	nums:=[]int{4,1,2,1,2}
	fmt.Println(singleNumber(nums))
}
