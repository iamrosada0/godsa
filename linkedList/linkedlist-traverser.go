package linkedlist

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	tail   *node
	length int
}

func (l linkedlist) Traverser() {
	// Começa a partir do head
	n := l.head

	// Loop para percorrer todos os nós
	for n != nil {
		// Exibe o valor do nó atual
		fmt.Printf("%d ", n.data)

		// Avança para o próximo nó
		n = n.next
	}

	// Exibe uma nova linha após a impressão
	fmt.Println()
}

func main() {
	// Exemplo de uso
	list := linkedlist{}
	list.head = &node{data: 1}
	list.head.next = &node{data: 2}
	list.head.next.next = &node{data: 3}
	list.tail = list.head.next.next
	list.length = 3

	// Percorre e exibe a lista
	list.Traverser() // Saída: 1 2 3
}
