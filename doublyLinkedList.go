package gods

import (
	"fmt"
)

type DoublyLinkedList[T comparable] struct {
	head *DblNode[T]
}

type DblNode[T comparable] struct {
	val  T
	next *DblNode[T]
	prev *DblNode[T]
}

func (list *DoublyLinkedList[T]) String() string {
	s := list.ToSlice()
	return fmt.Sprintf("%v", s)
}

func (list DoublyLinkedList[T]) Size() int {
	s := 0
	node := list.head
	for node != nil {
		s++
		node = node.next
	}
	return s
}

func (list *DoublyLinkedList[T]) Add(val T) {
	node := list.head
	if node == nil {
		list.head = &DblNode[T]{val: val, prev: list.head}
		return
	}
	for node.next != nil {
		node = node.next
	}
	node.next = &DblNode[T]{val: val, prev: node}
}

func (list *DoublyLinkedList[T]) Remove(val T) {
	if list.head == nil {
		return
	}
	if list.head.val == val {
		list.head = list.head.next
		list.head.prev = nil
		return
	}
	node := list.head
	for node != nil {
		if node.val == val {
			node.prev.next = node.next
			if node.next != nil {
				node.next.prev = node.prev
			}
			break
		}
		node = node.next
	}
}

func (list *DoublyLinkedList[T]) RemoveAll(val T) {
	if list.head == nil {
		return
	}
	for list.head.val == val {
		list.head = list.head.next
		list.head.prev = nil
	}
	node := list.head
	for node != nil {
		if node.val == val {
			node.prev.next = node.next
			if node.next != nil {
				node.next.prev = node.prev
			}
		}
		node = node.next
	}

}

func (list *DoublyLinkedList[T]) RemoveIf(checker func(T) bool) {
	if list.head == nil {
		return
	}
	for checker(list.head.val) {
		list.head = list.head.next
		list.head.prev = nil
	}
	node := list.head
	for node != nil {
		if checker(node.val) {
			node.prev.next = node.next
			if node.next != nil {
				node.next.prev = node.prev
			}
		}
		node = node.next
	}
}

func (list *DoublyLinkedList[T]) Contains(val T) bool {
	if list.head == nil {
		return false
	}
	node := list.head
	for node != nil {
		if node.val == val {
			return true
		}
		node = node.next
	}
	return false
}

func (list *DoublyLinkedList[T]) Poll() (T, bool) {
	if list.head == nil {
		return *new(T), false
	}
	if list.head.next == nil {
		val := list.head.val
		list.head = list.head.next
		return val, true
	}

	node := list.head
	for node.next != nil {
		node = node.next
	}
	val := node.val
	node.prev.next = nil
	return val, true
}

func (list *DoublyLinkedList[T]) ToSlice() []T {
	s := make([]T, 0, list.Size())
	node := list.head
	for node != nil {
		s = append(s, node.val)
		node = node.next
	}
	return s
}
