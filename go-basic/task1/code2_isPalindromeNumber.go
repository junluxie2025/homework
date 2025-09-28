package main

import "strconv"

// 回文数
func isPalindromeNumber(num int) bool {
	str := strconv.Itoa(num)
	strLen := len(str)
	for index, _ := range str {
		if str[index] != str[strLen-index-1] {
			return false
		}
	}
	return true
}
