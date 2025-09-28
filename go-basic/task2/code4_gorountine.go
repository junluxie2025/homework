package main

import (
	"fmt"
	"sync"
	"time"
)

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
type Task func()

type TaskResult struct {
	duration time.Duration
	id       int
}

type Schedule struct {
	tasks []Task
	wg    sync.WaitGroup
}

func statistic() {

	s := &Schedule{
		tasks: make([]Task, 0),
	}

	s.tasks = append(s.tasks, func() {
		fmt.Println("任务1开始")
		time.Sleep(1 * time.Second)
		fmt.Println("任务1结束")
	})

	s.tasks = append(s.tasks, func() {
		fmt.Println("任务2开始")
		time.Sleep(3 * time.Second)
		fmt.Println("任务2结束")
	})

	s.tasks = append(s.tasks, func() {
		fmt.Println("任务3开始")
		time.Sleep(5 * time.Second)
		fmt.Println("任务3结束")
	})

	s.wg.Add(len(s.tasks))

	ch := make(chan TaskResult, len(s.tasks))

	for i, task := range s.tasks {
		go func(id int, task Task) {
			defer s.wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			ch <- TaskResult{id: i, duration: duration}
		}(i, task)
	}

	s.wg.Wait()
	close(ch)

	for rs := range ch {
		fmt.Println(rs.id, " --> ", rs.duration)
	}
}
