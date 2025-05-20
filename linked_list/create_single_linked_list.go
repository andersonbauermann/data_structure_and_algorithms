package linked_list

import "fmt"

type SingleNode struct {
	Value int
	Next  *SingleNode
}

type SinglyLinkedList struct {
	Head *SingleNode
}

func (sll *SinglyLinkedList) AddToFront(value int) {
	newNode := &SingleNode{Value: value}
	newNode.Next = sll.Head
	sll.Head = newNode
}

func (sll *SinglyLinkedList) AddToEnd(value int) {
	newNode := &SingleNode{Value: value}
	if sll.Head == nil {
		sll.Head = newNode
	} else {
		current := sll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (sll *SinglyLinkedList) RemoveFromFront() *int {
	if sll.Head == nil {
		return nil
	}
	removedValue := sll.Head.Value
	sll.Head = sll.Head.Next
	return &removedValue
}

func PrintCreateSinglyLinkedList() {
	list := SinglyLinkedList{}

	list.AddToFront(10)
	list.AddToEnd(20)
	list.AddToFront(5)

	fmt.Println(*list.RemoveFromFront()) // 5
	fmt.Println(*list.RemoveFromFront()) // 10
	fmt.Println(*list.RemoveFromFront()) // 20
	fmt.Println(list.RemoveFromFront())  // nil
}
