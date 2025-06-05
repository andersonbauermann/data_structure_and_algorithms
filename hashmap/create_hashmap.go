package hashmap

import (
	"fmt"
)

type KeyValue struct {
	key   interface{}
	value interface{}
}

type HashMap struct {
	size    int
	buckets [][]KeyValue
}

func NewHashMap(size int) *HashMap {
	buckets := make([][]KeyValue, size)
	return &HashMap{
		size:    size,
		buckets: buckets,
	}
}

func (h *HashMap) hash(key interface{}) int {
	return int(fmt.Sprintf("%v", key)[0]) % h.size
}

func (h *HashMap) Put(key, value interface{}) {
	index := h.hash(key)
	bucket := h.buckets[index]

	for i, kv := range bucket {
		if kv.key == key {
			h.buckets[index][i].value = value
			return
		}
	}

	h.buckets[index] = append(bucket, KeyValue{key, value})
}

func (h *HashMap) Get(key interface{}) interface{} {
	index := h.hash(key)
	bucket := h.buckets[index]

	for _, kv := range bucket {
		if kv.key == key {
			return kv.value
		}
	}

	return nil
}

func (h *HashMap) Remove(key interface{}) {
	index := h.hash(key)
	bucket := h.buckets[index]

	for i, kv := range bucket {
		if kv.key == key {
			h.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

func (h *HashMap) String() string {
	result := ""

	for _, bucket := range h.buckets {
		if len(bucket) > 0 {
			for _, kv := range bucket {
				result += fmt.Sprintf("%v: %v\n", kv.key, kv.value)
			}
		}
	}

	return result
}

func PrintHashMap() {
	hashMap := NewHashMap(100)

	hashMap.Put("anderson", 30)
	hashMap.Put("naila", 25)
	hashMap.Put(60, []interface{}{"cachorro", 13})

	fmt.Println(hashMap)

	fmt.Println("anderson:", hashMap.Get("anderson"))

	hashMap.Remove("anderson")

	fmt.Println("anderson after remove:", hashMap.Get("anderson"))
}

/*
Estrutura:
	- Usa um slice de buckets: [][]KeyValue
	- Cada bucket é uma lista (slice) de pares chave-valor
	- Cada chave é mapeada para um índice com base em uma função hash

Função hash:
	- Recebe uma chave (interface{})
	- Converte a chave em string e usa o primeiro caractere como base
	- Retorna um índice: int(fmt.Sprintf("%v", key)[0]) % size
	Obs: Função simples, apenas para fins didáticos — pode causar colisões com frequência

Métodos principais:

1. Put(key, value)
	- Insere ou atualiza um par chave-valor
	- Etapas:
		- Calcula o índice via hash(key)
		- Busca no bucket se a chave já existe
			- Se sim, atualiza o valor
			- Se não, adiciona o novo par ao final do bucket

2. Get(key)
	- Retorna o valor associado à chave
	- Etapas:
		- Calcula o índice via hash(key)
		- Percorre o bucket procurando a chave
		- Retorna o valor se encontrar, senão retorna nil

3. Remove(key)
	- Remove o par com a chave informada
	- Etapas:
		- Calcula o índice via hash(key)
		- Percorre o bucket e remove o par se encontrar a chave

4. String()
	- Retorna uma string com todos os pares chave-valor do mapa
	- Percorre todos os buckets e concatena os pares encontrados

Observações:
	- Usa interface{} para aceitar qualquer tipo como chave e valor
	- O tratamento de colisão é feito por encadeamento (lista em cada bucket)
	- Código educacional: a função hash é simplificada
	- Pode ser melhorado com hashing mais eficiente e suporte a generics
*/
