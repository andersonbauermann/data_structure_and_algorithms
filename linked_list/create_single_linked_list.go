package linked_list

import "fmt"

type SingleNode struct {
	value int
	next  *SingleNode
}

type SinglyLinkedList struct {
	head *SingleNode
}

func (sll *SinglyLinkedList) AddToFront(value int) {
	newNode := &SingleNode{value: value}
	newNode.next = sll.head
	sll.head = newNode
}

func (sll *SinglyLinkedList) AddToEnd(value int) {
	newNode := &SingleNode{value: value}
	if sll.head == nil {
		sll.head = newNode
	} else {
		current := sll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (sll *SinglyLinkedList) RemoveFromFront() *int {
	if sll.head == nil {
		return nil
	}
	removedValue := sll.head.value
	sll.head = sll.head.next
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
