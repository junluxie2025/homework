package main

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func addOne(arr []int) (res []int) {

	length := len(arr)
	temp := make([]int, length+1)
	temp[length] = 1
	for i := length - 1; i >= 0; i-- {
		temp[i+1] = arr[i] + temp[i+1]
		yushu := temp[i+1] % 10
		beishu := temp[i+1] / 10
		temp[i+1] = yushu
		temp[i] = beishu
	}
	if temp[0] == 0 {
		return temp[1 : length+1]
	} else {
		return temp
	}
}
