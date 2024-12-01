package binarysearchtree

/*
Understanding the Algorithm
Main Insertion Function (Insert):

The first step checks if the root node of the tree is null (empty tree).
If the tree is empty, the new value becomes the root.
If the tree is not empty, the function delegates the

task to a helper function (InsertNode),
starting from the root node.
Helper Function (InsertNode):

This function recursively navigates the tree,
comparing the value to be inserted with the current node.
Depending on the comparison result,
it moves to either the left or right subtree.
Recursive Logic:

If the value is less than the current node's value,
move to the left.

If the value is greater than or equal to the current node's value, move to the right.
If the correct position is found (i.e., the left or right child is null), insert the value there.
*/
type Node struct {
	Value int
	Left  *Node
	Right *Node
}
type BST struct {
	Root *Node
}

func (bst *BST) Insert(value int) {

	if bst.Root == nil {
		bst.Root = &Node{Value: value}
	} else {
		insertNode(bst.Root, value)

	}

}

func insertNode(node *Node, value int) {

	if value < node.Value {
		//ESQUERDA
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			insertNode(node.Left, value)
		}
	} else {
		// DIRETA
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			insertNode(node.Right, value)
		}
	}

}
