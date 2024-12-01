package binarysearchtree

import "fmt"

// Node represents a node in the binary search tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BST represents a binary search tree.
type BST struct {
	Root *Node
}

// Remove deletes a value from the BST and returns whether the deletion was successful.
func (bst *BST) Remove(value int) bool {
	if bst.Root == nil {
		// If the tree is empty, there's nothing to remove
		return false
	}

	// Find the node to remove and its parent node
	nodeToRemove, parent := findNodeAndParent(bst.Root, nil, value)
	if nodeToRemove == nil {
		// If the value is not found in the tree, return false
		return false
	}

	// Case 1: Node to remove is a leaf node (no children)
	if nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		if parent == nil {
			// If the node to remove is the root and the only node in the tree
			bst.Root = nil
		} else if parent.Left == nodeToRemove {
			// If the node to remove is a left child
			parent.Left = nil
		} else {
			// If the node to remove is a right child
			parent.Right = nil
		}

	} else if nodeToRemove.Left == nil {
		// Case 2: Node to remove has only a right subtree
		if parent == nil {
			// If the node to remove is the root, make the right subtree the new root
			bst.Root = nodeToRemove.Right
		} else if parent.Left == nodeToRemove {
			// If the node to remove is a left child, link its parent's left pointer to the right subtree
			parent.Left = nodeToRemove.Right
		} else {
			// If the node to remove is a right child, link its parent's right pointer to the right subtree
			parent.Right = nodeToRemove.Right
		}
	} else if nodeToRemove.Right == nil {
		// Case 3: Node to remove has only a left subtree
		if parent == nil {
			// If the node to remove is the root, make the left subtree the new root
			bst.Root = nodeToRemove.Left
		} else if parent.Left == nodeToRemove {
			// If the node to remove is a left child, link its parent's left pointer to the left subtree
			parent.Left = nodeToRemove.Left
		} else {
			// If the node to remove is a right child, link its parent's right pointer to the left subtree
			parent.Right = nodeToRemove.Left
		}
	} else {
		// Case 4: Node to remove has both left and right subtrees
		// Find the largest value in the left subtree (in-order predecessor)
		largestNode, largestParent := findLargestNode(nodeToRemove.Left, nodeToRemove)

		// Replace the value of the node to remove with the largest value
		nodeToRemove.Value = largestNode.Value

		// Remove the largest node from the left subtree
		if largestParent.Left == largestNode {
			// If the largest node is a left child, remove it by linking to its left subtree
			largestParent.Left = largestNode.Left
		} else {
			// If the largest node is a right child, link the parent's right pointer to its left subtree
			largestParent.Right = largestNode.Left
		}
	}

	// Successfully removed the node
	return true
}

// Helper function to find a node and its parent
func findNodeAndParent(node, parent *Node, value int) (*Node, *Node) {
	if node == nil {
		// Base case: reached a null node, value not found
		return nil, nil
	}

	if value == node.Value {
		// Found the node with the matching value
		return node, parent
	}

	if value < node.Value {
		// Recur into the left subtree if the value is smaller
		return findNodeAndParent(node.Left, node, value)
	}

	// Recur into the right subtree if the value is larger
	return findNodeAndParent(node.Right, node, value)
}

// Helper function to find the largest node in a subtree and its parent
func findLargestNode(node, parent *Node) (*Node, *Node) {
	for node.Right != nil {
		// Traverse the right subtree until the largest node is found
		parent = node
		node = node.Right
	}
	return node, parent
}

// Main function to test deletion
func main() {
	bst := &BST{}

	// Test deletion of various cases (populate the BST first in a real test)
	fmt.Println("Delete 7:", bst.Remove(7))   // Leaf node
	fmt.Println("Delete 15:", bst.Remove(15)) // Node with one child
	fmt.Println("Delete 10:", bst.Remove(10)) // Node with two children

	fmt.Println("In-order traversal after deletion:")
}
