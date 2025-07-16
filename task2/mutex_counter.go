package main
import (
	"fmt"
	"sync"
)

type Counter struct{
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment(){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func main(){
	var counter Counter
	var wg sync.WaitGroup

	for i:=0;i<10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			for j :=0;j<1000;j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("计数器最终值：",counter.Value())
}