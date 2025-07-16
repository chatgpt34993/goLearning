package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

type AtomicCounter struct {
    value int64 // 使用int64类型以支持原子操作
}

// Increment 原子地增加计数器值
func (c *AtomicCounter) Increment() {
    atomic.AddInt64(&c.value, 1)
}

// Value 原子地获取计数器当前值
func (c *AtomicCounter) Value() int64 {
    return atomic.LoadInt64(&c.value)
}

func main() {
    var counter AtomicCounter
    var wg sync.WaitGroup

    // 启动10个协程
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // 每个协程递增1000次
            for j := 0; j < 1000; j++ {
                counter.Increment()
            }
        }()
    }

    wg.Wait() // 等待所有协程完成
    fmt.Println("计数器最终值:", counter.Value())
}    