package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Uso: go run main.go [binary_search|reverse_string]")
		return
	}

	switch os.Args[1] {
	case "binary_search":
		printBinarySearch()
	case "reverse_string":
		printReverse()
	default:
		println("Opção inválida.'")
	}
}
