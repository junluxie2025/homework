package main

import (
	"fmt"
	"sync"
)

func OnlyReceive(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Printf("收到数据: %d \n", i)
	}
}

func OnlySend(ch chan<- int, wg *sync.WaitGroup, count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		ch <- i
	}
	close(ch)
}

func res7() {
	w := sync.WaitGroup{}
	ch := make(chan int)
	w.Add(1)
	go OnlySend(ch, &w, 10)
	w.Add(1)
	go OnlyReceive(ch, &w)
	w.Wait()
}
