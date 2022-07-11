package gods

import "fmt"

type SinglyLinkedList[T comparable] struct {
	head *Node[T]
}

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

func (list *SinglyLinkedList[T]) String() string {
	s := make([]T, 0, list.Size())
	node := list.head
	for {
		if node == nil {
			break
		}
		s = append(s, node.val)
		node = node.next
	}
	return fmt.Sprintf("%v", s)
}

func (list SinglyLinkedList[T]) Size() int {
	s := 0
	node := list.head
	for {
		if node == nil {
			break
		}
		s++
		node = node.next
	}
	return s
}

func (list *SinglyLinkedList[T]) Add(val T) {
	node := list.head
	if node == nil {
		list.head = &Node[T]{val: val}
		return
	}
	for {
		if node.next == nil {
			node.next = &Node[T]{val: val}
			break
		}
		node = node.next
	}
}

func (list *SinglyLinkedList[T]) Remove(val T) {
	if list.head == nil {
		return
	}
	if list.head.val == val {
		list.head = list.head.next
		return
	}
	node := list.head
	for {
		if node.next == nil {
			break
		}
		if node.next.val == val {
			node.next = node.next.next
			break
		}
		node = node.next
	}
}

func (list *SinglyLinkedList[T]) RemoveAll(val T) {
	if list.head == nil {
		return
	}
	for list.head.val == val {
		list.head = list.head.next
	}
	node := list.head
	for {
		if node.next == nil {
			break
		}
		if node.next.val == val {
			node.next = node.next.next
		} else {
			node = node.next
		}
	}

}

func (list *SinglyLinkedList[T]) RemoveIf(checker func(T) bool) {
	if list.head == nil {
		return
	}
	for checker(list.head.val) {
		list.head = list.head.next
	}
	node := list.head
	for {
		if node.next == nil {
			break
		}
		if checker(node.next.val) {
			node.next = node.next.next
		} else {
			node = node.next
		}
	}
}

func (list *SinglyLinkedList[T]) Contains(val T) bool {
	if list.head == nil {
		return false
	}
	node := list.head
	for {
		if node == nil {
			break
		}
		if node.val == val {
			return true
		}
		node = node.next
	}
	return false
}
