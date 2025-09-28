package main

import (
	"time"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printOddAndEven() {
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				print(i, "\t")
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				print(i, "\t")
			}
		}
	}()
	time.Sleep(5 * time.Second)
}
