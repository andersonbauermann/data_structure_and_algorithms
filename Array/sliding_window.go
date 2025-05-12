package main

import "fmt"

func maxLengthSubstring(s string) int {
	l, r := 0, 0
	maxLen := 1
	counter := make(map[byte]int)

	counter[s[0]] = 1

	for r < len(s)-1 {
		r++

		if count, exists := counter[s[r]]; exists {
			counter[s[r]] = count + 1
		} else {
			counter[s[r]] = 1
		}

		for counter[s[r]] == 3 {
			counter[s[l]]--
			l++
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen
}

func printSlidingWindow() {
	str := "bcbbbcbaab"
	maxArray := maxLengthSubstring(str)
	fmt.Println("String testada:", str)
	fmt.Println("Maior string:", maxArray)
}
