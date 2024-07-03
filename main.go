package main

func firstUniqChar(s string) int {
	if len(s) <= 0 {
		return -1
	}

	dic := make(map[rune]bool)
	strArr := []rune(s)
	for i := 0; i < len(s); i++ {
		temp := strArr[i]
		if _, ok := dic[temp]; ok { // 不唯一
			dic[temp] = false
		} else {
			dic[temp] = true
		}
	}

	for i := 0; i < len(strArr); i++ {
		unique := dic[strArr[i]]
		if unique {
			return i
		}
	}

	return -1
}
