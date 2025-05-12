package main

import (
	"fmt"
)

func reverseWords(s string) string {
	chars := []rune(s)
	start := 0

	for end := 0; end <= len(chars); end++ {
		if end == len(chars) || chars[end] == ' ' {
			reverseSegment(chars, start, end-1)
			start = end + 1
		}
	}

	return string(chars)
}

func reverseSegment(chars []rune, start, end int) {
	for start < end {
		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}
}

func printReverse() {
	fmt.Println(reverseWords("hello world"))
	fmt.Println(reverseWords("Go is awesome"))
}
