package graph

type Graph map[string]map[string]int

var graph = Graph{
	"A": {"B": 1, "C": 4},
	"B": {"A": 1, "C": 2, "D": 5},
	"C": {"A": 4, "B": 2, "D": 1},
	"D": {"B": 5, "C": 1},
}

func dijkstra(graph Graph, start string) map[string]int {
	distances := make(map[string]int)
	visited := make(map[string]bool)

	for node := range graph {
		distances[node] = int(^uint(0) >> 1)
	}
	distances[start] = 0

	for len(visited) < len(graph) {
		minNode := ""
		minDistance := int(^uint(0) >> 1)

		for node, distance := range distances {
			if !visited[node] && distance < minDistance {
				minDistance = distance
				minNode = node
			}
		}

		if minNode == "" {
			break
		}

		visited[minNode] = true

		for neighbor, weight := range graph[minNode] {
			if !visited[neighbor] {
				newDistance := distances[minNode] + weight
				if newDistance < distances[neighbor] {
					distances[neighbor] = newDistance
				}
			}
		}
	}

	return distances
}

func PrintShortestPath(start string) {
	distances := dijkstra(graph, start)
	for node, distance := range distances {
		if distance == int(^uint(0)>>1) {
			println("Distance from", start, "to", node, "is infinity")
		} else {
			println("Distance from", start, "to", node, "is", distance)
		}
	}
}

/*
Objetivo:
Calcular a menor distância do vértice inicial (start) até todos os outros vértices de um grafo.

Conceito:
Dijkstra é um algoritmo de busca de caminhos mínimos.
Funciona apenas com pesos positivos.

Estrutura:
- graph → grafo representado por mapa: cada vértice aponta para seus vizinhos com pesos.
- start → vértice inicial.
- distances → armazena a menor distância conhecida até cada nó.
- visited → marca os nós já processados.

Comportamento passo a passo:
1. Inicializa:
   - Todas as distâncias como infinito (int(^uint(0) >> 1)).
   - A distância do start é 0.

2. Enquanto ainda houver nós não visitados:
   - Escolhe o nó com menor distância conhecida que ainda não foi visitado (minNode).
   - Marca esse nó como visitado.
   - Para cada vizinho do minNode:
     - Calcula nova distância: distância do minNode + peso da aresta.
     - Se for menor que a distância atual conhecida do vizinho, atualiza.

Resultado:
Retorna o mapa "distances" com as menores distâncias do start até todos os nós do grafo.

Usos comuns:
- Navegação GPS (caminho mais curto).
- Redes de computadores (roteamento de pacotes).
- Jogos (IA encontrando o caminho mais curto).
*/
