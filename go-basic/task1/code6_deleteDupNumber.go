package main

// 给你一无序序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度
func deleteDupNumber(arr []int) int {
	length := len(arr)
	i := 0
	k := length - 1
	for index := range length {
		i = index
		if i == k {
			break
		}
		for j := k; j > i; j-- {
			if arr[i] == arr[j] {
				//交换数据
				temp := arr[j]
				arr[j] = arr[k]
				arr[k] = temp
				k--
			}
		}
	}

	return i + 1
}
