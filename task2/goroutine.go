package main

import (
	"fmt"
	"sync"
)

//奇数
func printOdd(wg  *sync.WaitGroup){
	defer wg.Done()
    for i :=1;i<=10;i+=2 {
		fmt.Println("奇数：",i)
	}

}
//偶数
func printEven(wg *sync.WaitGroup){
	defer wg.Done()
	for i:=2; i<=10; i+=2 {
		fmt.Println("偶数：",i)
	}
}
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd(&wg)
	go printEven(&wg)

	wg.Wait()

}