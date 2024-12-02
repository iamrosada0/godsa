package heaps

import (
	"fmt"
)

// MaxHeap define uma estrutura para armazenar os elementos do heap máximo.
type MaxHeap struct {
	array []int // Array que representa o heap.
}

// Insert insere um valor no heap.
func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value) // Adiciona o valor no final do array.
	h.maxHeapifyUp(len(h.array) - 1) // Ajusta a posição do novo valor para manter a propriedade do MaxHeap.
}

// parent retorna o índice do nó pai de um índice fornecido.
func parent(index int) int {
	return (index - 1) / 2 // Fórmula para obter o índice do pai no heap.
}

// left retorna o índice do filho esquerdo de um índice fornecido.
func left(index int) int {
	return 2*index + 1 // Fórmula para obter o índice do filho esquerdo.
}

// right retorna o índice do filho direito de um índice fornecido.
func right(index int) int {
	return 2*index + 2 // Fórmula para obter o índice do filho direito.
}

// swap troca os valores de dois índices no array do heap.
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1] // Troca os valores nos índices especificados.
}

// maxHeapifyUp ajusta o heap de baixo para cima após a inserção para garantir a propriedade do MaxHeap.
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] { // Enquanto o pai for menor que o nó atual:
		h.swap(parent(index), index) // Troca o nó atual com o pai.
		index = parent(index)        // Atualiza o índice para continuar verificando mais acima na árvore.
	}
}

// findIndex encontra o índice de um valor no heap. Retorna -1 se o valor não for encontrado.
func (h *MaxHeap) findIndex(value int) int {
	for i, v := range h.array { // Itera sobre o array do heap.
		if v == value { // Se encontrar o valor:
			return i // Retorna o índice.
		}
	}
	return -1 // Retorna -1 se o valor não estiver no heap.
}

// Remove remove um valor do heap e ajusta a estrutura para manter a propriedade do MaxHeap.
func (h *MaxHeap) Remove(value int) bool {
	index := h.findIndex(value) // Encontra o índice do valor a ser removido.
	if index == -1 {            // Se o valor não estiver no heap:
		return false // Retorna falso indicando falha.
	}

	lastIndex := len(h.array) - 1       // Obtém o índice do último elemento.
	h.array[index] = h.array[lastIndex] // Substitui o valor removido pelo último elemento.
	h.array = h.array[:lastIndex]       // Remove o último elemento do array.

	h.maxHeapifyDown(index) // Ajusta a estrutura do heap de cima para baixo.
	return true             // Retorna verdadeiro indicando sucesso.
}

// maxHeapifyDown ajusta o heap de cima para baixo para manter a propriedade do MaxHeap após uma remoção.
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1 // Índice do último elemento no array.
	leftChild := left(index)      // Índice do filho esquerdo do nó atual.
	rightChild := right(index)    // Índice do filho direito do nó atual.

	for leftChild <= lastIndex { // Enquanto o filho esquerdo existir:
		largerChild := leftChild // Assume inicialmente que o filho esquerdo é o maior.

		if rightChild <= lastIndex && h.array[rightChild] > h.array[leftChild] { // Se o filho direito existir e for maior que o esquerdo:
			largerChild = rightChild // O filho maior agora é o direito.
		}

		if h.array[index] >= h.array[largerChild] { // Se o nó atual for maior ou igual ao maior filho:
			break // A propriedade do MaxHeap está preservada; pode sair do loop.
		}

		h.swap(index, largerChild) // Troca o nó atual com o maior filho.
		index = largerChild        // Atualiza o índice para continuar ajustando na próxima iteração.
		leftChild = left(index)    // Atualiza o índice do filho esquerdo.
		rightChild = right(index)  // Atualiza o índice do filho direito.
	}
}

// Função main para testar o MaxHeap.
func main() {
	m := &MaxHeap{} // Inicializa um novo MaxHeap vazio.

	fmt.Println(m) // Exibe o estado inicial do heap.

	heapsBuild := []int{12, 3, 5, 10, 20, 30} // Array de valores para construir o heap.

	for _, v := range heapsBuild { // Itera sobre cada valor no array.
		m.Insert(v) // Insere o valor no heap.
	}

	fmt.Println(m.array) // Exibe o estado final do heap.
}
