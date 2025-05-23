package linked_list

import "fmt"

func hasLoop(head *SingleNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func PrintVerifyLoop() {
	list := SinglyLinkedList{}

	list.AddToFront(10)
	list.AddToEnd(20)
	list.AddToFront(5)
	list.AddToEnd(30)

	list.Head.Next.Next.Next = list.Head // CRIA O LOOP

	if hasLoop(list.Head) {
		fmt.Println("The linked list has a loop.")
	} else {
		fmt.Println("The linked list does not have a loop.")
	}
}

/*
As variáveis -fast- e -slow- recebem o head
a variável -fast- avança duas posições a cada iteração e a variável -slow -avança uma posição a cada iteração
Se -fast- chegar ao final da linkedList, quer dizer que não existe loop na lista, mas se houver algum loop, as variáveis -fast- e -slow- serão iguais em alguma iteração, retornando o valor true
*/
