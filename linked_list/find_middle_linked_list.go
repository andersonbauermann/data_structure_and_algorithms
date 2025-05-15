package linked_list

import "fmt"

func findMiddle(head *SingleNode) *SingleNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func PrintFindMiddle() {
	list := SinglyLinkedList{}

	list.AddToFront(10)
	list.AddToEnd(20)
	list.AddToFront(5)
	list.AddToEnd(30)

	fmt.Print("Original List: ")
	printList(list.head)

	middleNode := findMiddle(list.head)

	if middleNode != nil {
		fmt.Printf("Middle Node: %d\n", middleNode.value)
	} else {
		fmt.Println("List is empty.")
	}
}

// as variáveis slow e fast são usadas para percorrer a lista.
// O ponteiro slow avança um nó por vez, enquanto o ponteiro fast avança dois nós por vez.
// dessa forma, quando o ponteiro fast chegar no final da linkedList, o ponteiro slow vai estar exatamente no meio da linkedList
