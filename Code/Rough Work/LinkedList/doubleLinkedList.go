package main

import "fmt"

//Node1 struct
type Node1 struct {
	property int
	nextNode *Node1
	prevNode *Node1
}

//DoubleLinkedList struct
type DoubleLinkedList struct {
	header *Node1
}

//AddHeader func
func (doubleLinkedList *DoubleLinkedList) AddHeader(property int) {

	node := &Node1{}
	node.property = property
	node.nextNode = nil
	node.prevNode = nil

	if doubleLinkedList.header != nil {

		node.nextNode = doubleLinkedList.header
		doubleLinkedList.header.prevNode = node
	}
	doubleLinkedList.header = node
}

//Add func
func (doubleLinkedList *DoubleLinkedList) Add(property int) {

	node := &Node1{}
	node.property = property
	node.nextNode = nil
	node.prevNode = nil

	temp := doubleLinkedList.FindLastNode()
	if temp != nil {
		temp.nextNode = node
		node.prevNode = temp
	}
}

//FindLastNode func
func (doubleLinkedList *DoubleLinkedList) FindLastNode() *Node1 {

	var node *Node1

	for node = doubleLinkedList.header; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			break
		}
	}
	if node != nil {
		return node
	}
	return nil
}

//AddAt func
func (doubleLinkedList *DoubleLinkedList) AddAt(property int) {

	node := &Node1{}
	node.property = property
	node.nextNode = nil
	node.prevNode = nil

	var temp *Node1

	for temp = doubleLinkedList.header; temp != nil; temp = node.nextNode {
		if temp.property == property {
			node.nextNode = temp.nextNode
			node.prevNode = temp
			temp.nextNode = node
			break
		}
	}
}

//Iterate func
func (doubleLinkedList *DoubleLinkedList) Iterate() {

	var temp *Node1

	for temp = doubleLinkedList.header; temp != nil; temp = temp.nextNode {
		fmt.Println(temp.property)
	}
}

func main() {

	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.AddHeader(1)
	doubleLinkedList.Add(2)
	doubleLinkedList.Add(3)
	doubleLinkedList.Add(4)
	doubleLinkedList.Add(5)
	doubleLinkedList.Iterate()
}
