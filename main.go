package main

import (
	"fmt"
)

//custom error
type listEmpty struct{}

func (e *listEmpty) Error() string {
	return "List is empty"
}

//ListFeature interface
type ListFeature interface {
	add(value int) List
	delete(value int) bool
	length() int
	update(valueToUpdate int, newValue int) bool
}

type Node struct {
	data int
	next *Node
	prev *Node
}
type List struct {
	head, tail *Node
}

func (l *List) First() *Node {
	return l.head
}
func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}

func (l List) add(value int) List {
	n := &Node{data: value}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n
	return l

}

func (l List) delete(value int) bool {
	deleted := false

	for n := l.First(); n != nil && !deleted; n = n.Next() {
		if n.data == value {
			fmt.Println("Deleting Node ")
			prev_node := n.prev
			next_node := n.next
			// Remove this node
			prev_node.next = n.next
			next_node.prev = n.prev
			deleted = true
		}
	}
	return deleted
}

func display(headNode *Node) error {

	if headNode == nil {
		return &listEmpty{}
	} else {
		fmt.Println(headNode.data)
		display(headNode.next)
	}
	return nil

}

func (l List) length() int {

	node := l.head
	counter := 0
	for node == nil {
		node = node.next
		counter += 1
	}

	return counter
}

func (l List) update(valueToUpdate int, newValue int) bool {
	updated := false

	for n := l.First(); n != nil && !updated; n = n.Next() {
		if n.data == valueToUpdate {
			fmt.Println("Updating Node ")
			n.data = newValue
			updated = true
		}
	}
	return updated
}

func main() {
	var list List
	var listFeat ListFeature = list

	//adding elements
	listFeat = listFeat.add(2)
	listFeat = listFeat.add(3)
	listFeat = listFeat.add(4)
	listFeat = listFeat.add(6)

	//delete specific node from the list
	listFeat.delete(3)

	//passing head of list to display
	display(list.head)

}
