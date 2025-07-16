package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 表示一个可执行的任务
type Task struct {
	Name      string        // 任务名称
	Execute   func()        // 任务执行函数
	TimeTaken time.Duration // 执行耗时
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks []*Task
	wg    sync.WaitGroup
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks: make([]*Task, 0),
	}
}

// AddTask 添加任务到调度器
func (s *TaskScheduler) AddTask(name string, task func()) {
	s.tasks = append(s.tasks, &Task{
		Name:    name,
		Execute: task,
	})
}

// Run 并发执行所有任务
func (s *TaskScheduler) Run() {
	s.wg.Add(len(s.tasks))

	for _, task := range s.tasks {
		go func(t *Task) {
			defer s.wg.Done()
			
			start := time.Now()
			t.Execute()              // 执行任务
			t.TimeTaken = time.Since(start)
			
			fmt.Printf("任务 %s 执行完成，耗时: %v\n", t.Name, t.TimeTaken)
		}(task)
	}

	s.wg.Wait() // 等待所有任务完成
}

func main() {
	scheduler := NewTaskScheduler()

	// 添加任务
	scheduler.AddTask("任务A", func() {
		time.Sleep(2 * time.Second) // 模拟耗时操作
	})

	scheduler.AddTask("任务B", func() {
		time.Sleep(1 * time.Second) // 模拟耗时操作
	})

	scheduler.AddTask("任务C", func() {
		time.Sleep(1500 * time.Millisecond) // 模拟耗时操作
	})

	// 执行所有任务
	scheduler.Run()
}