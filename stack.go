package immutable

// Stack implements a last in, first out container. Nil and the zero value for Stack are both empty
// stacks.
type Stack struct {
	top    interface{}
	bottom *Stack
}

// Returns true if the stack is empty.
//
// Complexity: O(1)
func (s *Stack) Empty() bool {
	return s == nil || s.bottom == nil
}

// Returns the top item on the stack.
//
// Complexity: O(1)
func (s *Stack) Peek() interface{} {
	return s.top
}

// Removes the top item from the stack.
//
// Complexity: O(1)
func (s *Stack) Pop() *Stack {
	return s.bottom
}

// Places an item onto the top of the stack.
//
// Complexity: O(1)
func (s *Stack) Push(value interface{}) *Stack {
	return &Stack{
		top:    value,
		bottom: s,
	}
}
