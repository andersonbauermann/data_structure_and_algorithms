package main

import (
	"fmt"
	"math/rand"
)

func search(nums []int, target int) (int, int) {
	left, right, steps := 0, len(nums)-1, 0

	for left <= right {
		steps++
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle, steps
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1, steps
}

func printBinarySearch() {
	nums := make([]int, 101)
	for i := 0; i <= 100; i++ {
		nums[i] = i
	}

	target := rand.Intn(101)

	fmt.Println("Array:", nums)
	fmt.Println("Elemento a ser encontrado:", target)

	result, steps := search(nums, target)

	if result != -1 {
		fmt.Printf("Elemento encontrado na posição: %d, número de passos executados %d", result, steps)
	} else {
		fmt.Println("Elemento não encontrado.")
	}
}
