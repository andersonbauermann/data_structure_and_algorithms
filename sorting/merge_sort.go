package sorting

import (
	"data_structure/linked_list"
)

func merge(node1 *linked_list.SingleNode, node2 *linked_list.SingleNode) *linked_list.SingleNode {
	head := &linked_list.SingleNode{}
	tail := head

	for node1 != nil && node2 != nil {
		if node1.Value < node2.Value {
			tail.Next = node1
			node1 = node1.Next
		} else {
			tail.Next = node2
			node2 = node2.Next
		}
		tail = tail.Next
	}

	if node1 != nil {
		tail.Next = node1
	} else {
		tail.Next = node2
	}

	return head.Next
}

func mergeSort(head *linked_list.SingleNode) *linked_list.SingleNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := linked_list.FindMiddle(head)
	after_middle := middle.Next
	middle.Next = nil
	left := mergeSort(head)
	right := mergeSort(after_middle)

	return merge(left, right)
}

func PrintMergeSort() {
	list := linked_list.SinglyLinkedList{}

	list.AddToFront(10)
	list.AddToEnd(20)
	list.AddToFront(5)
	list.AddToEnd(11)
	list.AddToFront(80)
	list.AddToEnd(2)
	list.AddToFront(7)
	list.AddToEnd(27)
	list.AddToFront(23)

	sorted := mergeSort(list.Head)
	linked_list.PrintList(sorted)
}
