package main

import (
	"sync"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printOddAndEven() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i += 2 {
			print(i, "\t")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i += 2 {
			print(i, "\t")
		}
	}()
	wg.Wait()
}
