package main

import (
	"time"
)

func res8() {
	ch := make(chan int, 4)
	go OnlySend(ch)
	go OnlyReceive(ch)
	time.Sleep(1 * time.Second)
}
