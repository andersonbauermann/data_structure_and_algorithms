package linked_list

import "fmt"

func FindMiddle(head *SingleNode) *SingleNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
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
	PrintList(list.Head)

	middleNode := FindMiddle(list.Head)

	if middleNode != nil {
		fmt.Printf("Middle Node: %d\n", middleNode.Value)
	} else {
		fmt.Println("List is empty.")
	}
}

// as variáveis slow e fast são usadas para percorrer a lista.
// O ponteiro slow avança um nó por vez, enquanto o ponteiro fast avança dois nós por vez.
// dessa forma, quando o ponteiro fast chegar no final da linkedList, o ponteiro slow vai estar exatamente no meio da linkedList
