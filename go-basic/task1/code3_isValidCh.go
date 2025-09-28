package main

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValidCh(str string) bool {
	kv := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	items := []string{}

	for _, v := range str {
		item := string(v)
		items = append(items, item)
		temp := kv[item]
		if temp != "" {
			if len(items)-2 < 0 {
				return false
			} else if temp != items[len(items)-2] {
				return false
			} else {
				items = items[:len(items)-1-1]
			}
		}

	}

	return len(items) == 0
}
