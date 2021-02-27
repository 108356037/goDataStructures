package main

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/arraystack"
)

type Node struct {
	value int
	next  *Node
}

type linkedList struct {
	listName string
	head     *Node
}

func (list *linkedList) traverse() error {
	currentNode := list.head
	fmt.Printf("%+v, at mem address %p\n", *currentNode, currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v, at mem address %p\n", *currentNode, currentNode)
	}
	return nil
}

func createLinkedList(s string) *linkedList {
	return &linkedList{
		listName: s,
		head:     nil,
	}
}

func (list *linkedList) addNodeTail(val int) error {
	newNode := &Node{
		value: val,
		next:  nil,
	}
	if list.head == nil {
		list.head = newNode
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	return nil
}

func (list *linkedList) addNodeHead(val int) error {
	newNode := &Node{
		value: val,
		next:  nil,
	}
	if list.head == nil {
		list.head = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
	return nil
}

func (list *linkedList) addNodeAtGivenVal(newNodeVal, searchVal int) error {
	newNode := &Node{
		value: newNodeVal,
		next:  nil,
	}
	if list.head == nil {
		list.head = newNode
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			if currentNode.value == searchVal {
				newNode.next = currentNode.next
				currentNode.next = newNode
				return nil
			} else {
				currentNode = currentNode.next
			}
		}
		currentNode.next = newNode
	}
	return nil
}

func (list *linkedList) reverseListIteration() error {
	fmt.Println("Reversing list...")
	currentNode := list.head
	var prev *Node = nil
	var next *Node = nil
	for currentNode != nil {
		next = currentNode.next
		currentNode.next = prev
		prev = currentNode
		currentNode = next
	}
	list.head = prev
	return nil
}

func (list *linkedList) reverseListStack() *linkedList {
	stack := arraystack.New()
	currentNode := list.head
	for currentNode != nil {
		stack.Push(currentNode.value)
		currentNode = currentNode.next
	}
	newList := createLinkedList("Reversedlist")
	for val := range stack.Values() {
		newList.addNodeTail(val)
	}
	return newList
}

func main() {

	my := createLinkedList("v1")
	my.addNodeTail(20)
	my.addNodeTail(10)
	my.addNodeHead(40)
	my.addNodeHead(50)
	my.addNodeAtGivenVal(100, 15)
	my.addNodeAtGivenVal(101, 13)
	my.addNodeAtGivenVal(144, 100)
	my.addNodeAtGivenVal(144, 50)
	my.traverse()

	// rv := my.reverseListStack()
	// rv.traverse()
	my.reverseListIteration()
	my.traverse()

}
