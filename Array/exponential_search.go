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

/*
Um algoritmo de busca para encontrar um elemento em um array ordenado.
Combina expansão rápida (exponencial) + busca binária refinada.

Funciona em duas etapas:
	-Fase de expansão: Dobra o índice de busca até achar um valor maior ou igual ao alvo.
	-Fase de refinamento: Usa busca binária no intervalo encontrado.

Melhor que busca binária pura se o elemento estiver no começo.
Mais eficaz em arrays muito grandes (diminui o espaço de busca rapidamente).
Complexidade:
	O(log n) para a fase de expansão.
	O(log n) para a fase de refinamento.
*/
