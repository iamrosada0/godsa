package linkedlist

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

	n := l.Head

	for n != nil && n.Data != value {

		n = n.Next
	}

	if n == nil {
		return false
	}

	return true

}
