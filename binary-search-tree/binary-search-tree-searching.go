package binarysearchtree

/*
1. the root = ∅ in which case value is not in the BST; or

2. root.Value = value in which case value is in the BST; or

3. value < root.Value, we must inspect the left subtree of
root for value; or

4. value > root.Value, we must inspect
the right subtree of root for value.
*/
type Node struct {
	Value int
	Right *Node
	Left  *Node
}

type BST struct {
	Root *Node
}

// Search checks if a value exists in the BST
func (n BST) Search(value int) bool {
	if n.Root == nil {
		// If the tree is empty, the value cannot exist
		return false
	}

	// Start searching from the root
	return Contains(n.Root, value)
}

// Contains is a recursive function that searches for the value in the tree
func Contains(node *Node, value int) bool {
	if node == nil {
		// Base case: reached a leaf node and did not find the value
		return false
	}

	if node.Value == value {
		// Base case: value found
		return true
	}

	if value < node.Value {
		// Recurse into the left subtree if the value is smaller
		return Contains(node.Left, value)
	}

	// Recurse into the right subtree if the value is larger
	return Contains(node.Right, value)
}
