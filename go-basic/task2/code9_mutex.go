package main

import (
	"fmt"
	"sync"
	"time"
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
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				safeCount.Increment()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("计数器的值: %d", safeCount.GetCount())
}
