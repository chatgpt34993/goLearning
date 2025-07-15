package main

import "fmt"

func isPalindrome(x int) bool {
if x<0 {
	return false
}
original := x
    reversed := 0
    for x != 0 {
        reversed = reversed*10 + x%10
        x /= 10
    }
    return original == reversed

}


func main(){
	testCases := []int{121, -121, 10, 12321}
    for _, num := range testCases {
        fmt.Printf("输入: %d, 输出: %v\n", num, isPalindrome(num))
    }

}
