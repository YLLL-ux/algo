package main

func isSubsequence(s string, t string) bool {
	x := len(s)
	y := len(t)
	if x == 0 && y == 0 {
		return true
	}

	i := 0
	j := 0
	for i < x && j < y {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == x
}
