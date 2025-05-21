package binary_tree

// DFS = Depth First Search - Busca em profundidade

import "fmt"

func (bt *BinaryTree) DFS(data int) bool {
	return dfsRecursive(bt.Root, data)
}

func dfsRecursive(node *Node, data int) bool {
	if node != nil {
		fmt.Println(node.Value)
	} else {
		return false
	}

	if node.Value == data {
		return true
	}

	if dfsRecursive(node.Left, data) {
		return true
	}
	if dfsRecursive(node.Right, data) {
		return true
	}

	return false
}

func PrintDFS() {
	binaryTree := BinaryTree{}
	valuesToInsert := []int{10, 5, 15, 3, 7, 12, 18}

	for _, value := range valuesToInsert {
		binaryTree.Root = Insert(binaryTree.Root, value)
	}

	binaryTree.DFS(18)
}

/*
DFS (Depth-First Search), ou Busca em Profundidade, é um algoritmo que explora o caminho o mais fundo possível antes de voltar.
Ou seja, ele desce pelos filhos de um nó até não conseguir mais, e só então volta para tentar outros caminhos.

Passos do algoritmo recursivo:
1. Verifica se o nó atual é nulo:
   - Se for, retorna false (não encontrou).
2. Verifica se o valor do nó é igual ao que procuramos:
   - Se for, retorna true.
3. Busca na subárvore esquerda.
4. Se não encontrou, busca na subárvore direita.
5. Se não encontrou em nenhum dos dois, retorna false.

Vantagens do DFS:
- Simples de implementar com recursão.
- Usa menos memória que BFS.
- Útil para:
  - Verificar se um valor está presente.
  - Percorrer todos os nós.
  - Resolver labirintos, expressões, etc.

Exemplo em linguagem natural:
Imagine que você está em um labirinto e sempre decide:
- "Primeiro tento ir pela esquerda até o fim."
- "Se não der certo, volto e tento pela direita."
*/
