package main

import "fmt"

func exponentialSearch(arr []int, target int) (int, int) {
	steps := 1

	if arr[0] == target {
		return 0, steps
	}

	n := len(arr)
	i := 1
	for i < n && arr[i] < target {
		steps++
		i *= 2
	}

	if arr[i] == target {
		steps++
		return i, steps
	}

	left, right := i/2, min(i, n-1) // esse min(..) é para evitar o erro de index out of range

	return search(arr, target, &left, &right, &steps)
}

func printExponentialSearch() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 8

	result, steps := exponentialSearch(arr, target)

	if result != -1 {
		fmt.Printf("Elemento encontrado na posição: %d, número de passos executados %d", result, steps)
	} else {
		println("Elemento não encontrado.")
	}
}
