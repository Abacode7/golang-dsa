package queue

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return &Queue{data: make([]interface{}, 0)}
}

func (r *Queue) Enqueue(value interface{}) {
	r.data = append(r.data, value)
}

func (r *Queue) Dequeue() interface{} {
	if len(r.data) == 0 {
		return nil
	}
	output := r.data[0]
	r.data = r.data[1:]
	return output
}

func (r *Queue) Peek() interface{} {
	if len(r.data) == 0 {
		return nil
	}
	return r.data[0]
}
