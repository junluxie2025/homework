package main

import (
	"fmt"
	"time"
)

func OnlyReceive(ch <-chan int) {
	for i := range ch {
		fmt.Printf("收到数据: %d \n", i)
	}
}

func OnlySend(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("发送数据: %d \n", i)
	}
	close(ch)
}

func rse7() {
	ch := make(chan int, 1)
	go OnlySend(ch)
	go OnlyReceive(ch)
	time.Sleep(1 * time.Second)
}
