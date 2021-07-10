package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (t *BinarySearchTree) Insert(value int) {
	node := &Node{value: value}
	if t.root == nil {
		t.root = node
		return
	}

	currentNode := t.root
	for currentNode != nil {
		if value < currentNode.value {
			if currentNode.left != nil {
				currentNode = currentNode.left
			} else {
				currentNode.left = node
				return
			}
		} else {
			if currentNode.right != nil {
				currentNode = currentNode.right
			} else {
				currentNode.right = node
				return
			}
		}
	}
}

func (t *BinarySearchTree) Lookup(value int) *Node {
	if t.root == nil {
		return nil
	}

	currentNode := t.root
	for currentNode != nil {
		if value < currentNode.value {
			currentNode = currentNode.left
		} else if value > currentNode.value {
			currentNode = currentNode.right
		} else {
			return currentNode
		}
	}

	return nil
}

func (t *BinarySearchTree) Remove(value int) {
	if t.root == nil {
		return
	}

	var parentNode *Node
	currentNode := t.root
	for currentNode != nil {
		if value < currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.left
		} else if value > currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil && currentNode.right == nil {
				if parentNode == nil {
					t.root = nil
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = nil
					} else {
						parentNode.right = nil
					}
				}
			} else if currentNode.right == nil {
				if parentNode == nil {
					t.root = currentNode.left
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = currentNode.left
					} else {
						parentNode.right = currentNode.left
					}
				}
			} else if currentNode.right.left == nil {
				if parentNode == nil {
					t.root = currentNode.right
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = currentNode.right
					} else {
						parentNode.right = currentNode.right
					}
				}
			} else {
				leftMost := currentNode.right.left
				leftMostParent := currentNode.right
				for leftMost.left != nil {
					leftMostParent = leftMost
					leftMost = leftMost.left
				}

				leftMostParent.left = leftMost.right
				leftMost.left = currentNode.left
				leftMost.right = currentNode.right
				if parentNode == nil {
					t.root = leftMost
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = leftMost
					} else {
						parentNode.right = leftMost
					}
				}
			}
			return
		}
	}
}

func (t *BinarySearchTree) Traverse(node *Node) {
	fmt.Println(node.value)
	if node.left != nil {
		t.Traverse(node.left)
	}
	if node.right != nil {
		t.Traverse(node.right)
	}
}

func main() {
	tree := BinarySearchTree{}
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)
	tree.Remove(20)
	tree.Traverse(tree.root)
	// fmt.Println(tree.Lookup(20))
	// fmt.Println(tree.Lookup(30))
}
