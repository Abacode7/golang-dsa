package stack

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	data := make([]interface{}, 0)
	return &Stack{data: data}
}

/** Todo: push, pop, peek*/

// Push adds value to the (conceptual) head of the stack
// O(1) time, O(1) extra space
// NB: The append operation extra time O(N) when the underlying array allocation
// is full and moves value to a one with a bigger capacity
func (r *Stack) Push(item interface{}) {
	r.data = append(r.data, item)
}

// Pop removes value from the (conceptual) head of the stack
// O(1) time, O(1) extra space
// NB: The extract operator (:) still references the same underlying array
func (r *Stack) Pop() interface{} {
	if len(r.data) == 0 {
		return nil
	}

	lastItem := r.data[len(r.data)-1]
	r.data = r.data[:len(r.data)-1]
	return lastItem
}

// Peek views value from the (conceptual) head of the stack
// O(1) time, O(1) extra space
func (r *Stack) Peek() interface{} {
	if len(r.data) == 0 {
		return nil
	}
	return r.data[len(r.data)-1]
}
