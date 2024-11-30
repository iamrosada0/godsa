package linkedlist

import "fmt"

//create a loop and inside of loop we need to check if exist the value

type Node struct {
	Data int
	Next *Node
}

type Linkedlist struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *Linkedlist) Search(value int) bool {
	current := l.Head

	for current != nil {
		if current.Data == value {
			return true // Value found
		}
		current = current.Next
	}

	return false // Value not found
}

func main() {
	list := Linkedlist{}

	// Adding some nodes
	list.Head = &Node{Data: 10}
	list.Head.Next = &Node{Data: 20}
	list.Head.Next.Next = &Node{Data: 30}

	// Searching for values
	fmt.Println(list.Search(20)) // Output: true
	fmt.Println(list.Search(40)) // Output: false
}
