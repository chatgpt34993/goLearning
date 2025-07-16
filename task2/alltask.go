package main

import "fmt"

func increaseByTen(numPtr  *int) {
	*numPtr +=10
}

func doubleSlice(slicePtr *[]int){
	slice :=*slicePtr
	for i:=range slice {
		slice[i] *=2
	}

}

func main(){
	num :=5
	fmt.Println("修改前的值：",num)

	increaseByTen(&num)
	fmt.Println("修改后的值：",num)

	numbers :=[]int{1,2,3,4,5}
   fmt.Println("修改前的切片:", numbers)

   doubleSlice(&numbers)

   fmt.Println("修改后的切片:", numbers)

}