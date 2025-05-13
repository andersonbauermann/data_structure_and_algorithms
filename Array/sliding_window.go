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

/*
Um algoritmo de busca eficiente para encontrar padrões ou somas em subconjuntos contínuos de um array
Funciona através de uma “janela” que se move sobre o array, mantendo o controle de uma parte do array por vez, sem repetir cálculos desnecessários

Como funciona:
	-Define o tamanho da janela (quantos elementos serão analisados de cada vez)
	-Calcula o valor inicial da janela (por exemplo, a soma dos primeiros k elementos)
	-Desliza a janela uma posição por vez:
	-Remove o elemento que saiu da janela
	-Adiciona o novo elemento que entrou
	-Atualiza o valor (por exemplo, soma) conforme necessário
	-Repete até percorrer todo o array ou encontrar a condição desejada

Complexidade:
O(n) no pior caso
*/
