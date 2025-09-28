package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func res10() {
	w := sync.WaitGroup{}
	w.Add(10)
	cnt := atomic.Int32{}
	for i := 0; i < 10; i++ {
		go func() {
			defer w.Done()
			for j := 0; j < 1000; j++ {
				cnt.Add(1)
			}
		}()
	}
	w.Wait()
	fmt.Printf("计数器的值: %d", cnt.Load())
}
