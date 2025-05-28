package heap

import (
	"errors"
	"fmt"
)

type MinHeap struct {
	heap []int
}

func (h *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) leftChild(index int) int {
	return 2*index + 1
}

func (h *MinHeap) rightChild(index int) int {
	return 2*index + 2
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 && h.heap[h.parent(index)] > h.heap[index] {
		h.heap[h.parent(index)], h.heap[index] = h.heap[index], h.heap[h.parent(index)]
		index = h.parent(index)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	smallest := index
	left := h.leftChild(index)
	right := h.rightChild(index)

	if left < len(h.heap) && h.heap[left] < h.heap[smallest] {
		smallest = left
	}

	if right < len(h.heap) && h.heap[right] < h.heap[smallest] {
		smallest = right
	}

	if smallest != index {
		h.heap[index], h.heap[smallest] = h.heap[smallest], h.heap[index]
		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.heapifyUp(len(h.heap) - 1)
}

func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.heap) == 0 {
		return 0, errors.New("heap is empty")
	}

	if len(h.heap) == 1 {
		min := h.heap[0]
		h.heap = []int{}
		return min, nil
	}

	root := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapifyDown(0)

	return root, nil
}

func (h *MinHeap) Peek() (int, error) {
	if len(h.heap) == 0 {
		return 0, errors.New("heap is empty")
	}
	return h.heap[0], nil
}

func (h *MinHeap) Size() int {
	return len(h.heap)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.heap) == 0
}

func PrintMinHeap() {
	h := &MinHeap{}
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	fmt.Println("Min:", h.heap)

	min, _ := h.ExtractMin()
	fmt.Println("Extracted Min:", min)

	fmt.Println("After Extraction:", h.heap)
}

/*
MinHeap (Heap Mínimo)
Tipo de estrutura: Árvore binária completa (implementada como array)
Propriedade:
    - Todo nó é menor que seus filhos.
    - O menor valor está sempre na raiz (índice 0 do array).
Exemplo (representação como árvore):
         1
        / \
       3   8
      /
     5

Array correspondente:
    [1, 3, 8, 5]

Comportamento:
    - A cada inserção, o valor é adicionado no final do array e depois sobe (heapify up) até estar na posição correta.
    - Ao remover o menor elemento (raiz), o último valor do array vai para o topo e é ajustado para baixo (heapify down).

Usos comuns:
    - Algoritmos de caminho mínimo (ex: Dijkstra)
    - Gerenciamento de prioridades (ex: filas de prioridade)
    - Simulações e agendamento de tarefas

---

Método: ExtractMin()
Objetivo:
    Remove e retorna o menor valor da heap (a raiz).

Funcionamento:
    - Se a heap estiver vazia, retorna erro.
    - Se houver apenas um elemento, remove e retorna esse único valor.
    - Caso contrário:
        1. Guarda o valor da raiz (heap[0]).
        2. Move o último valor do array para a posição 0 (topo da heap).
        3. Remove o último valor do array (com pop).
        4. Chama o heapifyDown para reorganizar a heap.
    - Retorna o valor original da raiz (o menor).

Exemplo:
    Heap antes: [1, 3, 8, 5]
    - Remove 1 (menor valor).
    - Move 5 para o topo.
    - Após reorganizar: [3, 5, 8]
    Resultado: retorna 1

Importância:
    Essencial para manter a propriedade do heap após remoção.
    Usado quando queremos sempre acessar o menor valor de forma eficiente (O(log n)).
*/
