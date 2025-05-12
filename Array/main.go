package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Uso: go run main.go [binary_search|reverse_string|sliding_window|exponential_search]")
		return
	}

	switch os.Args[1] {
	case "binary_search":
		printBinarySearch()
	case "reverse_string":
		printReverse()
	case "sliding_window":
		printSlidingWindow()
	case "exponential_search":
		printExponentialSearch()
	default:
		println("Opção inválida.'")
	}
}
