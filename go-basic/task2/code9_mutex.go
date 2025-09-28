package main

import (
	"fmt"
	"sync"
)

type SafeCount struct {
	mutex sync.Mutex
	count int
}

func (s *SafeCount) Increment() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.count++
}

func (s *SafeCount) GetCount() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.count
}

func res9() {
	safeCount := SafeCount{count: 0}
	w := sync.WaitGroup{}
	w.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer w.Done()
			for j := 0; j < 1000; j++ {
				safeCount.Increment()
			}
		}()
	}
	w.Wait()
	fmt.Printf("计数器的值: %d", safeCount.GetCount())
}
