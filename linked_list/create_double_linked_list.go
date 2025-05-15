package linked_list

import "fmt"

type DoubleNode struct {
	value int
	next  *DoubleNode
	prev  *DoubleNode
}

type DoublyLinkedList struct {
	head *DoubleNode
	tail *DoubleNode
}

func (reference *DoublyLinkedList) AddToFront(value int) {
	newNode := &DoubleNode{value: value}
	if reference.head == nil {
		reference.head = newNode
		reference.tail = newNode
	} else {
		newNode.next = reference.head
		reference.head.prev = newNode
		reference.head = newNode
	}
}

func (reference *DoublyLinkedList) AddToEnd(value int) {
	newNode := &DoubleNode{value: value}
	if reference.tail == nil {
		reference.head = newNode
		reference.tail = newNode
	} else {
		newNode.prev = reference.tail
		reference.tail.next = newNode
		reference.tail = newNode
	}
}

func (reference *DoublyLinkedList) RemoveFromFront() *int {
	if reference.head == nil {
		return nil
	}
	removedValue := reference.head.value
	if reference.head == reference.tail {
		reference.head = nil
		reference.tail = nil
	} else {
		reference.head = reference.head.next
		reference.head.prev = nil
	}
	return &removedValue
}

func (reference *DoublyLinkedList) RemoveFromEnd() *int {
	if reference.tail == nil {
		return nil
	}
	removedValue := reference.tail.value
	if reference.head == reference.tail {
		reference.head = nil
		reference.tail = nil
	} else {
		reference.tail = reference.tail.prev
		reference.tail.next = nil
	}
	return &removedValue
}

func PrintCreateDoublyLinkedList() {
	list := DoublyLinkedList{}

	list.AddToFront(10)
	list.AddToEnd(20)
	list.AddToFront(5)

	fmt.Println(*list.RemoveFromFront()) // 5
	fmt.Println(*list.RemoveFromEnd())   // 20
	fmt.Println(*list.RemoveFromFront()) // 10
	fmt.Println(list.RemoveFromFront())  // nil
}
