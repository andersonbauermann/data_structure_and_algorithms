package binary_tree

import "fmt"

func preorderTraversal(node *Node, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Value)
	preorderTraversal(node.Left, result)
	preorderTraversal(node.Right, result)
}

func inorderTraversal(node *Node, result *[]int) {
	if node == nil {
		return
	}
	inorderTraversal(node.Left, result)
	*result = append(*result, node.Value)
	inorderTraversal(node.Right, result)
}

func postorderTraversal(node *Node, result *[]int) {
	if node == nil {
		return
	}
	postorderTraversal(node.Left, result)
	postorderTraversal(node.Right, result)
	*result = append(*result, node.Value)
}

func PrintTraversals() {
	binaryTree := BinaryTree{}
	valuesToInsert := []int{10, 5, 15, 3, 7, 12, 18}

	for _, value := range valuesToInsert {
		binaryTree.Root = Insert(binaryTree.Root, value)
	}

	preorderResult := []int{}
	inorderResult := []int{}
	postorderResult := []int{}

	preorderTraversal(binaryTree.Root, &preorderResult)
	inorderTraversal(binaryTree.Root, &inorderResult)
	postorderTraversal(binaryTree.Root, &postorderResult)

	fmt.Println("Pre-order Traversal:", preorderResult)
	fmt.Println("In-order Traversal:", inorderResult)
	fmt.Println("Post-order Traversal:", postorderResult)
}

/*
Postorder (Pós-ordem)
Ordem: Esquerda → Direita → Raiz
Comportamento:
	-Visita primeiro todos os nós da subárvore esquerda.
	-Depois os da direita.
	-E por último, o nó atual (raiz).
Usos comuns:
Útil para deletar a árvore (liberar memória).
Também usado em algumas operações de avaliação de expressão (ex: árvores de expressão matemática).

Preorder (Pré-ordem)
Ordem: Raiz → Esquerda → Direita
Comportamento:
	-Visita primeiro o nó atual (raiz).
	-Depois visita os nós da subárvore esquerda.
	-Por fim, os da subárvore direita.
Usos comuns:
Útil para copiar uma árvore.
Serve para salvar a estrutura da árvore (por exemplo, serialização).

Inorder (Em ordem)
Ordem: Esquerda → Raiz → Direita
Comportamento:
	-Visita todos os nós da subárvore esquerda.
	-Depois visita o nó atual (raiz).
	-Depois visita todos os nós da subárvore direita.
Resultado:
Em uma árvore binária de busca (BST), retorna os valores em ordem crescente.
*/
