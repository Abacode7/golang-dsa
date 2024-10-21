package queue

import "testing"

func TestNewQueue(t *testing.T) {
	queue := NewQueue()

	if queue == nil {
		t.Error("Expected valid queue, not nil")
	}

	if queue.data == nil {
		t.Error("Expected queue to have empty data, not nil")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("Apple")
	queue.Enqueue(1)
	queue.Enqueue("Zebra")

	if queue.Peek() != "Apple" {
		t.Error("Expected value at the head of queue be Apple, but got: ", queue.Dequeue())
	}

	if queue.Peek() != 1 {
		t.Error("Expected value at the head of queue be 1, but got: ", queue.Dequeue())
	}

	if queue.Peek() != "Zebra" {
		t.Error("Expected value at the head of queue be Zebra, but got: ", queue.Dequeue())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("Cats")
	queue.Enqueue(7)
	queue.Enqueue("Dogs")

	value1 := queue.Dequeue()
	if value1 != "Apple" {
		t.Error("Expected value at the head of queue be Cats, but got: ", value1)
	}

	value2 := queue.Dequeue()
	if value2 != 7 {
		t.Error("Expected value at the head of queue be 7, but got: ", value2)
	}

	value3 := queue.Dequeue()
	if value3 != "Dogs" {
		t.Error("Expected value at the head of queue be Dogs, but got: ", value3)
	}
}
