package main

import "fmt"

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func sliceOp(sli *[]int) {
	for i, v := range *sli {
		(*sli)[i] = v * 2
	}
}

func res2() {
	fmt.Print("修改后的值:")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceOp(&arr)
	for _, v := range arr {
		fmt.Print(v, "\t")
	}
	fmt.Println("")
}
