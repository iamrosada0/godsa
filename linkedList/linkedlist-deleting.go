package linkedlist

/*
Explanation of Each Scenario
The list is empty:

	If l.Head is nil, the list is empty, and the function immediately returns false.

The node to remove is the only node in the linked list:

	When l.Head is equal to l.Tail, and the value matches,
	the list is cleared by setting l.Head and l.Tail to nil and updating the length.

Removing the head node:

	If the value matches l.Head.Data, update l.Head to point to the next node.
	If the list becomes empty after this operation, also update l.Tail.

Removing the tail node:

	During traversal, if current matches the l.Tail, update l.Tail to prev (the previous node).

Removing a node between head and tail:

	Adjust the Next pointer of the prev node to skip the current node.

The item to remove doesn’t exist in the linked list:

	If the traversal completes without finding a match, return false.
*/
type node struct {
	data int   // O valor armazenado no nó
	next *node // Ponteiro para o próximo nó
}

type linkedlist struct {
	head   *node // Ponteiro para o primeiro nó da lista
	tail   *node // Ponteiro para o último nó da lista
	length int   // Número total de nós na lista
}

// Método para deletar um item da lista
func (l *linkedlist) deleteItem(value int) bool {

	// Caso 1: Lista vazia
	if l.head == nil {
		// Não há nada para deletar, retorna falso
		return false
	}

	// Caso 2: A lista tem apenas um nó e é o que queremos deletar
	if l.head == l.tail && l.head.data == value {
		l.head = nil // Remove o único nó, ajustando o head
		l.tail = nil // Ajusta o tail para nil também
		l.length = 0 // A lista agora está vazia, então o comprimento é 0
		return true  // Retorna verdadeiro porque o nó foi deletado
	}

	// Caso 3: O valor a ser deletado está no primeiro nó (head)
	if l.head.data == value {
		l.head = l.head.next // Avança o head para o próximo nó
		l.length--           // Decrementa o comprimento da lista

		// Se, ao mover o head, ele se tornar nil (lista vazia)
		if l.head == nil {
			l.tail = nil // Ajusta o tail também para nil
		}
		// O nó foi deletado, então retorna verdadeiro
		return true
	}

	// Caso 4: Percorrer a lista para encontrar o valor
	prev := l.head         // Aponta para o nó anterior ao nó atual
	current := l.head.next // Aponta para o próximo nó após o anterior

	// Loop para encontrar o nó com o valor desejado
	for current != nil {
		// Verifica se o nó atual contém o valor
		if current.data == value {
			// Caso 5: O nó a ser deletado é o último nó (tail)
			if current == l.tail {
				l.tail = prev // Atualiza o tail para o nó anterior
			}

			// Caso 6: O nó a ser deletado está no meio
			prev.next = current.next // Ajusta o ponteiro do nó anterior para "pular" o nó atual
			l.length--               // Decrementa o comprimento da lista
			return true              // Nó deletado com sucesso
		}

		// Avança para o próximo nó
		prev = current         // Atualiza o nó anterior para o atual
		current = current.next // Atualiza o nó atual para o próximo
	}

	// Caso 7: O valor não foi encontrado na lista
	return false
}

/*Explicação Geral:

Verifica a lista vazia:
Se l.head == nil, a lista está vazia,
então não há nada para remover.

Deleta o único nó da lista:
Se o nó a ser removido é o único nó (head == tail),
ajustamos head e tail para nil e zeramos o comprimento da lista.

Deleta o primeiro nó:
Se o valor a ser removido está no head,
avançamos head para o próximo nó e ajustamos tail caso o novo head seja nil.

Percorre a lista:
Usamos um loop para verificar cada nó,
com dois ponteiros:

prev aponta para o nó anterior.
current aponta para o nó atual.
Atualiza ponteiros para excluir o nó:
Se current contém o valor a ser deletado:

Atualizamos prev.next para pular o nó atual.
Ajustamos tail se o nó atual era o último.
Valor não encontrado:
Se o loop termina sem encontrar o valor, retornamos false.

*/
