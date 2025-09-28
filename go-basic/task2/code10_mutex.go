package main

import (
	"fmt"
	"sync/atomic"

	"time"
)

func res10() {
	cnt := atomic.Int32{}
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				cnt.Add(1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("计数器的值: %d", cnt.Load())
}
