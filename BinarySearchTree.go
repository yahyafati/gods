package gods

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BinarySearchTree[T constraints.Ordered] struct {
	root *BSTNode[T]
}

type BSTNode[T constraints.Ordered] struct {
	left   *BSTNode[T]
	right  *BSTNode[T]
	parent *BSTNode[T]
	val    T
}

func (node *BSTNode[T]) String() string {
	return fmt.Sprintf("%v", node.val)
}
func (node *BSTNode[T]) isEmpty() bool {
	return (node.left == nil && node.right == nil)
}

func (bst *BinarySearchTree[T]) Size() int {
	return len(bst.ToSlice())
}

func (bst *BinarySearchTree[T]) String() string {
	return fmt.Sprintf("%v", bst.ToSlice())
}

func (bst *BinarySearchTree[T]) ToSlice() []T {
	queue := []*BSTNode[T]{}
	queue = append(queue, bst.root)
	slice := []T{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == nil {
			continue
		}
		slice = append(slice, current.val)
		queue = append(queue, current.left, current.right)
	}
	return slice
}

func (bst *BinarySearchTree[T]) ToSliceDFS() []T {
	stack := []*BSTNode[T]{}
	stack = append(stack, bst.root)
	slice := []T{}
	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]
		if current == nil {
			continue
		}
		slice = append(slice, current.val)

		newStack := []*BSTNode[T]{current.left, current.right}
		stack = append(newStack, stack...)
	}
	return slice
}

func (bst *BinarySearchTree[T]) Add(val T) {
	if bst.root == nil {
		bst.root = &BSTNode[T]{val: val}
		return
	}
	node := bst.root
	for {
		if val > node.val {
			if node.right == nil {
				node.right = &BSTNode[T]{val: val, parent: node}
				break
			}
			node = node.right
		} else {
			if node.left == nil {
				node.left = &BSTNode[T]{val: val, parent: node}
				break
			}
			node = node.left
		}
	}
}

func find[T constraints.Ordered](node *BSTNode[T], val T) *BSTNode[T] {
	if node == nil {
		return nil
	}
	if val == node.val {
		return node
	} else if val < node.val {
		return find(node.left, val)
	} else {
		return find(node.right, val)
	}
}

func (bst *BinarySearchTree[T]) Contains(val T) bool {
	return find(bst.root, val) != nil
}

func findRightMostParent[T constraints.Ordered](node *BSTNode[T]) *BSTNode[T] {
	// fmt.Println(node, node.right, node.left)
	if node.isEmpty() {
		return node
	}
	if node.right != nil {
		return findRightMostParent(node.right)
	}
	return findRightMostParent(node.left)
}

func (bst *BinarySearchTree[T]) Remove(val T) bool {
	node := find(bst.root, val)
	if node == nil {
		return false
	}
	rMost := findRightMostParent(node)
	node.val = rMost.val
	if rMost == rMost.parent.left {
		rMost.parent.left = nil
	} else {
		rMost.parent.right = nil
	}
	return true
}

func (bst *BinarySearchTree[T]) RemoveAll(val T) {
	for bst.Remove(val) {
	}
}
