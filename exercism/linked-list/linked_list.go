package linkedlist

import "errors"

// ErrEmptyList error
var ErrEmptyList = errors.New("error empty list")

// Node struct
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// Next return next node
func (node *Node) Next() *Node {
	return node.next
}

// Prev return previous node
func (node *Node) Prev() *Node {
	return node.prev
}

// First return first node
func (node *Node) First() *Node {
	cur := node
	for cur != nil && cur.prev != nil {
		cur = cur.prev
	}
	return cur
}

// Last return last node
func (node *Node) Last() *Node {
	cur := node
	for cur != nil && cur.next != nil {
		cur = cur.next
	}
	return cur
}

// List struct
type List struct {
	first *Node
	last  *Node
}

// NewList construct a new list
func NewList(args ...interface{}) *List {
	if len(args) == 0 {
		return &List{
			first: nil,
			last:  nil,
		}
	}
	list := &List{first: &Node{Val: args[0], next: nil, prev: nil}}
	cur := list.first
	for i := 1; i < len(args); i++ {
		cur.next = &Node{
			Val:  args[i],
			next: nil,
			prev: cur,
		}
		cur = cur.next
	}
	list.last = cur
	return list
}

// First return the first node
func (list *List) First() *Node {
	return list.first
}

// Last return the last node
func (list *List) Last() *Node {
	return list.last
}

// PushBack add new node the end
func (list *List) PushBack(item interface{}) {
	node := &Node{
		Val:  item,
		next: nil,
		prev: list.last,
	}
	if list.last == nil {
		list.last = node
		list.first = node
	} else {
		list.last.next = node
		list.last = list.last.next
	}
}

// PopBack remove and return the last node
func (list *List) PopBack() (interface{}, error) {
	if list.last == nil {
		return nil, ErrEmptyList
	}
	tmp := list.last
	list.last = list.last.prev
	if list.last != nil {
		list.last.next = nil
	} else {
		list.first = nil
	}
	return tmp.Val, nil
}

// PushFront add new node to the beginning
func (list *List) PushFront(item interface{}) {
	newNode := &Node{
		Val:  item,
		next: list.first,
		prev: nil,
	}
	if list.first == nil {
		list.first = newNode
		list.last = newNode
	} else {
		list.first = newNode
		list.first.next.prev = list.first
	}
}

// PopFront remove and return the first node
func (list *List) PopFront() (interface{}, error) {
	if list.first == nil {
		return nil, ErrEmptyList
	}
	tmp := list.first
	list.first = list.first.next
	if list.first != nil {
		list.first.prev = nil
	} else {

		list.last = nil
	}
	return tmp.Val, nil
}

// Reverse the linked list
func (list *List) Reverse() {
	cur := list.last
	list.last, list.first = list.first, list.last
	for cur != nil {
		temp := cur.prev
		cur.prev, cur.next = cur.next, cur.prev
		cur = temp
	}
}
