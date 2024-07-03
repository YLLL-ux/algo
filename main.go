package main

import "fmt"

func main() {
	s := "bbbaaaba"
	t := "aaabbbba"

	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		tempS := s[i]
		tempT := t[i]

		if (s2t[tempS] > 0 && s2t[tempS] != tempT) || (t2s[tempT] > 0 && t2s[tempT] != tempS) {
			return false
		}

		s2t[tempS] = tempT
		t2s[tempT] = tempS
	}

	return true
}
