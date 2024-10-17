package stack

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	data := make([]interface{}, 0)
	return &Stack{data: data}
}

/** Todo: push, pop, peek*/

func (r *Stack) Push(item interface{}) {
	r.data = append(r.data, item)
}

func (r *Stack) Pop() interface{} {
	if len(r.data) == 0 {
		return nil
	}

	lastItem := r.data[len(r.data)-1]
	r.data = r.data[:len(r.data)-1]
	return lastItem
}

func (r *Stack) Peek() interface{} {
	if len(r.data) == 0 {
		return nil
	}
	return r.data[len(r.data)-1]
}
