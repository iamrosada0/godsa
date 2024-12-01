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

func (l *linkedlist) ReverseTraversal() {
	// Caso base: lista vazia
	if l.tail == nil {
		fmt.Println("A lista está vazia.")
		return
	}

	// Começa pelo último nó
	curr := l.tail

	// Enquanto o nó atual não for o head
	for curr != l.head {
		// Encontra o predecessor do nó atual
		prev := l.head
		for prev.next != curr {
			prev = prev.next
		}

		// Exibe o valor do nó atual
		fmt.Printf("%d ", curr.data)

		// Move para o nó anterior
		curr = prev
	}

	// Exibe o valor do head
	fmt.Printf("%d\n", l.head.data)
}

func main() {
	// Exemplo de uso
	list := linkedlist{}
	list.head = &node{data: 5}
	list.head.next = &node{data: 10}
	list.head.next.next = &node{data: 1}
	list.head.next.next.next = &node{data: 40}
	list.tail = list.head.next.next.next
	list.length = 4

	// Percorre a lista na ordem inversa
	list.ReverseTraversal() // Saída: 40 1 10 5
}
