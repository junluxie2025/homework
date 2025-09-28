package main

// 查找字符串数组中的最长公共前缀
func longestCommonPrefix(str1 string, str2 string) string {
	runes1 := []rune(str1)
	runes2 := []rune(str2)
	length := min(len(runes1), len(runes2))
	cnt := 0
	for l := 0; l < length; l++ {
		if runes1[l] == runes2[l] {
			cnt += 1
		} else {
			break
		}
	}
	return string(runes1[:cnt])
}
