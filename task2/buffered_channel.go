package main
import (
	"fmt"
	"sync"
)
func main() {
    ch := make(chan int, 10) // 创建容量为10的缓冲通道
    var wg sync.WaitGroup

    // 生产者协程
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 1; i <= 100; i++ {
            ch <- i
            fmt.Printf("生产者发送: %d\n", i)
        }
        close(ch) // 关闭通道
    }()
     // 消费者协程
    wg.Add(1)
    go func() {
        defer wg.Done()
        for num := range ch {
            fmt.Printf("消费者接收: %d\n", num)
        }
    }()

	wg.Wait()


}