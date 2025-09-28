package main

import (
	"sync"
)

func res8() {
	w := sync.WaitGroup{}
	ch := make(chan int, 100)
	w.Add(1)
	go OnlySend(ch, &w, 100)
	w.Add(1)
	go OnlyReceive(ch, &w)
	w.Wait()
}
