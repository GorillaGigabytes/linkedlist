package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) AddANodeAtTheEnd(value int) {
	newNode := &Node{data: value}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head

		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

func (list *LinkedList) ViewTheList() {
	if list.head == nil {
		fmt.Println("The list is empty")
		return
	}

	current := list.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}
	list.AddANodeAtTheEnd(5)
	list.AddANodeAtTheEnd(3)
	list.AddANodeAtTheEnd(1)

	list.ViewTheList()
}

// Thank you
