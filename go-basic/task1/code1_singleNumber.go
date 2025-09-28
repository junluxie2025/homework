package main

// 136 只出现一次的数字
func singleNumber(arr []int) int {
	cnt := make(map[int]int)

	for _, v := range arr {
		cnt[v] += 1
	}

	for k, v := range cnt {
		if v == 1 {
			return k
		}
	}
	return 0
}
