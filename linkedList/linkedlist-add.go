package linkedlist

type Node struct {
	Data int
	Next *Node
}

type Linkedlist struct {
	Head   *Node
	Length int
	Tail   *Node
}

func (l *Linkedlist) Add(value int) {

	n := &Node{Data: value}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}

}
