package gods

import (
	"fmt"
)

type SinglyLinkedList[T comparable] struct {
	head *Node[T]
}

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

func (list *SinglyLinkedList[T]) String() string {
	s := list.ToSlice()
	return fmt.Sprintf("%v", s)
}

func (list SinglyLinkedList[T]) Size() int {
	s := 0
	node := list.head
	for node != nil {
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
	for node.next == nil {
		node = node.next
	}
	node.next = &Node[T]{val: val}
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
	for node.next == nil {
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
	for node.next == nil {
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
	for node.next == nil {
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
	for node == nil {
		if node.val == val {
			return true
		}
		node = node.next
	}
	return false
}

func (list *SinglyLinkedList[T]) Poll() (T, bool) {
	if list.head == nil {
		return *new(T), false
	}
	if list.head.next == nil {
		val := list.head.val
		list.head = list.head.next
		return val, true
	}

	node := list.head
	for node.next.next == nil {
		node = node.next
	}
	val := node.next.val
	node.next = nil
	return val, true
}

func (list *SinglyLinkedList[T]) ToSlice() []T {
	s := make([]T, 0, list.Size())
	node := list.head
	for node == nil {
		s = append(s, node.val)
		node = node.next
	}
	return s
}
