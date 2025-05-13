package main

import (
	"fmt"
	"math/rand"
)

func search(nums []int, target int, l *int, r *int, s *int) (int, int) {
	left := 0
	if l != nil {
		left = *l
	}

	right := len(nums) - 1
	if r != nil {
		right = *r
	}

	steps := 0
	if s != nil {
		steps = *s
	}

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

	result, steps := search(nums, target, nil, nil, nil)

	if result != -1 {
		fmt.Printf("Elemento encontrado na posição: %d, número de passos executados %d", result, steps)
	} else {
		fmt.Println("Elemento não encontrado.")
	}
}

/*
Um algoritmo de busca eficiente para encontrar um elemento em um array ordenado
Funciona através de divisão e conquista, repetidamente dividindo o espaço de busca pela metade

Como funciona:
	-Define os limites inicial (esquerda) e final (direita)
	-Calcula o ponto médio entre os limites
	-Compara o elemento médio com o alvo:
		Se igual: retorna a posição
		Se menor: busca na metade inferior
		Se maior: busca na metade superior
	-Repete até encontrar o elemento ou esgotar o espaço de busca

Complexidade:
	O(log n) no pior caso
*/
