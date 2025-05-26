package stack

import (
	"errors"
	"fmt"
)

type Node struct {
	value any
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

// Push - adiciona um item ao topo da pilha.
func (s *Stack) Push(item any) {
	newNode := &Node{value: item, next: s.top}
	s.top = newNode
	s.size++
}

// Pop - remove o item do topo da pilha e o retorna. Se a pilha estiver vazia, retorna um erro.
func (s *Stack) Pop() (any, error) {
	if s.top == nil {
		return nil, errors.New("empty stack")
	}

	poppedNode := s.top
	s.top = poppedNode.next
	s.size--

	return poppedNode.value, nil
}

// Peek - retorna o item do topo da pilha sem removê-lo. Se a pilha estiver vazia, retorna um erro.
func (s *Stack) Peek() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("empty stack")
	}

	return s.top.value, nil
}

// Size - retorna o número de itens na pilha.
func (s *Stack) Size() int {
	return s.size
}

func PrintStack() {
	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Tamanho da pilha:", stack.Size())

	top, err := stack.Peek()
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Topo da pilha:", top)
	}

	for i := 0; i < 3; i++ {
		item, err := stack.Pop()
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Println("Removido da pilha:", item)
		}
	}

	fmt.Println("Tamanho da pilha após remoção:", stack.Size())
}
