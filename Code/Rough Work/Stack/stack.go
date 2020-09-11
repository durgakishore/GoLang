package main

import "fmt"

//Stack struct
type Stack struct {
	elements      []int
	elementscount int
}

//Push func
func (stack *Stack) Push(element int) {

	s := Stack{}
	s.elements = append(s.elements, element)
	stack.elements = append(s.elements, stack.elements[:stack.elementscount]...)

	stack.elementscount = stack.elementscount + 1
	fmt.Println("Print ", stack.elements)
}

//Pop func
func (stack *Stack) Pop() {

	if stack.elementscount > 1 {
		stack.elements = stack.elements[1:stack.elementscount]
	} else {
		stack.elements = nil
	}
	stack.elementscount = stack.elementscount - 1
	fmt.Println("Print ", stack.elements)
}

func main() {

	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
}
