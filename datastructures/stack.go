package datastructures

import "errors"

// Inspired by http://www.code2succeed.com/stack-implementation-in-golang/

type item struct {
	value int //value as interface type to hold any data type
	next  *item
}

// Stack is implemented as linked list: Each item has a pointer to the next item in the stack
// Example:
//   stack := new(Stack)
//   stack.Push(1)
type Stack struct {
	top  *item
	size int
}

// Len returns the size of the Stack
func (stack *Stack) Len() int {
	return stack.size
}

// Push adds a number to the top of the stack
func (stack *Stack) Push(value int) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

// Pop returns the topmost number from the stack (the number that was most recently added)
func (stack *Stack) Pop() (value int, err error) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}

	return -1, errors.New("Stack is empty")
}
