package main

func longestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}

	counter := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	res := 0
	odd := 0
	for _, count := range counter {
		if count%2 == 0 {
			res += count
		} else {
			res += count - 1
			odd = 1
		}
	}

	return res + odd
}
