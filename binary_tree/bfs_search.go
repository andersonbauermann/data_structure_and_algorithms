package binary_tree

// BFS = Breadth First Search - Busca em largura

import "fmt"

func bfs(root *Node, target int) bool {
	if root == nil {
		return false
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Value)

		if node.Value == target {
			return true
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return false
}

func PrintBFS() {
	binaryTree := BinaryTree{}
	valuesToInsert := []int{10, 5, 15, 3, 7, 12, 18}

	for _, value := range valuesToInsert {
		binaryTree.Root = Insert(binaryTree.Root, value)
	}

	bfs(binaryTree.Root, 18)
}

/*
BFS (Breadth-First Search), ou Busca em Largura, é um algoritmo que explora todos os nós no mesmo nível antes de descer para o próximo.
Ou seja, ele visita todos os filhos de um nó antes de visitar os netos, depois os bisnetos, e assim por diante.

Passos do algoritmo:
1. Verifica se a raiz é nula:
   - Se for, retorna false (não encontrou).
2. Cria uma fila (queue) e insere a raiz.
3. Enquanto a fila não estiver vazia:
   a. Remove o primeiro nó da fila.
   b. Verifica se o valor é o que estamos procurando:
      - Se for, retorna true.
   c. Adiciona os filhos (esquerda e direita) do nó atual na fila.
4. Se terminar a fila e não encontrar, retorna false.

Vantagens do BFS:
- Ideal para encontrar o caminho mais curto em grafos ou árvores não ponderadas.
- Explora os níveis mais próximos da raiz primeiro.
- Útil para:
  - Encontrar o menor caminho.
  - Verificar conexões entre nós em grafos.
  - Percorrer a árvore por níveis (level-order traversal).

Exemplo em linguagem natural:
Imagine que você está em um labirinto e sempre decide:
- "Primeiro exploro todos os caminhos de uma sala antes de avançar para as próximas."
- "Analiso os caminhos mais próximos antes de seguir para mais longe."
*/
