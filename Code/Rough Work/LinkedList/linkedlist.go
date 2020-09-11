package main

import (
	"fmt"
)

//Node struct
type Node struct {
	property int
	nextNode *Node
}

//LinkedList struct
type LinkedList struct {
	headNode *Node
}

//AddtoHead header
func (linkedList *LinkedList) AddtoHead(property int) {

	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = node
}

//Add func
func (linkedList *LinkedList) Add(property int) {

	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var temp *Node

	for temp = linkedList.headNode; temp != nil; temp = temp.nextNode {
		if temp.nextNode == nil {
			temp.nextNode = node
			break
		}
	}
}

//FindNode func
func (linkedList *LinkedList) FindNode(property int) *Node {

	var node *Node
	var temp *Node

	for temp = linkedList.headNode; node != nil; node = node.nextNode {
		if temp.property == property {
			node = temp
			break
		}
	}
	return node
}

//AddAt func
func (linkedList *LinkedList) AddAt(property int) {

	node := &Node{}
	node.property = property
	node.nextNode = nil

	temp := linkedList.FindNode(property)

	if temp != nil {
		node.nextNode = temp.nextNode
		temp.nextNode = node
	}
}

//IterateList func
func (linkedList *LinkedList) IterateList() {

	var node *Node

	for node = linkedList.headNode; node != nil; node = node.nextNode {

		fmt.Println(node.property)
	}
}

func main() {

	linkedList := LinkedList{}

	linkedList.AddtoHead(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(5)
	linkedList.IterateList()

}
