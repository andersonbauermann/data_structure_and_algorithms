package linked_list

import "fmt"

func invertList(head *SingleNode) *SingleNode {
	var prev *SingleNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func PrintInvertSingleLinkedList() {
	list := SinglyLinkedList{}

	list.AddToFront(10)
	list.AddToEnd(20)
	list.AddToFront(5)

	fmt.Print("Original List: ")
	PrintList(list.Head)

	invertedHead := invertList(list.Head)

	fmt.Print("Inverted List: ")
	PrintList(invertedHead)
}

func PrintList(head *SingleNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

// invertList inverte uma lista simplesmente encadeada.
// A ideia é "virar os ponteiros" para que o último nó vire o primeiro.
//
// Exemplo antes da inversão: 1 -> 2 -> 3 -> nil
// Exemplo depois da inversão: 3 -> 2 -> 1 -> nil
//
// Funcionamento passo a passo:
// 1. Cria um ponteiro 'prev' como nil (ele será o novo final da lista).
// 2. Usa 'current' para percorrer a lista original.
// 3. Em cada passo:
//    - Guarda o próximo nó em 'next'.
//    - Faz o current.next apontar para 'prev' (inverte a direção).
//    - Move 'prev' para o current.
//    - Move 'current' para o next.
// 4. Quando o loop termina, 'prev' aponta para o novo primeiro nó (a lista invertida).
