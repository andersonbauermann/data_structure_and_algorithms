package main

import (
	"data_structure/array"
	"data_structure/linked_list"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Uso: go run main.go [binary_search|reverse_string|sliding_window|exponential_search]")
		return
	}

	switch os.Args[1] {
	case "binary_search":
		array.PrintBinarySearch()
	case "reverse_string":
		array.PrintReverse()
	case "sliding_window":
		array.PrintSlidingWindow()
	case "exponential_search":
		array.PrintExponentialSearch()
	case "create_doubly_linked_list":
		linked_list.PrintCreateDoublyLinkedList()
	case "create_singly_linked_list":
		linked_list.PrintCreateSinglyLinkedList()
	case "invert_linked_list":
		linked_list.PrintInvertSingleLinkedList()
	case "find_middle_linked_list":
		linked_list.PrintFindMiddle()
	case "veify_loop_linked_list":
		linked_list.PrintVerifyLoop()
	default:
		println("Opção inválida.'")
	}
}
