package main

import "fmt"

func main(){
	ch :=make(chan int) //无缓冲通道

	//生产者协程
	go func(){
		for i:=1;i<=10;i++ {
			ch <- i
			fmt.Printf("已发送： %d\n",i)

		}
		close(ch)
	}()

	go func(){
		for num :=range ch{
         fmt.Printf("已接收： %d\n",num)
		}
	}()

	fmt.Scanln()

}