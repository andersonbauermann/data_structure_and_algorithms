package main

import (
	"data_structure/array"
	"data_structure/binary_tree"
	"data_structure/linked_list"
	"data_structure/sorting"
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
	case "bubble_sort":
		sorting.PrintBubbleSort()
	case "quick_sort":
		sorting.PrintQuickSort()
	case "merge_sort":
		sorting.PrintMergeSort()
	case "create_binary_tree":
		binary_tree.PrintBinaryTreeSearch()
	case "binary_tree_traversals":
		binary_tree.PrintTraversals()
	case "binary_tree_dfs":
		binary_tree.PrintDFS()
	case "binary_tree_bfs":
		binary_tree.PrintBFS()
	default:
		println("Opção inválida.'")
	}
}
