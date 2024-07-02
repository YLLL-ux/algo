package main

import "fmt"

func main() {
	res := isAnagram("abc", "abx")
	fmt.Println(res)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int, 0)
	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}

	return true
}
