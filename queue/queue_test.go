package queue

import (
	"testing"
)

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

	if queue.Dequeue() != "Apple" {
		t.Error("Expected value at the head of queue be Apple")
	}

	if queue.Dequeue() != 1 {
		t.Error("Expected value at the head of queue be 1")
	}

	if queue.Dequeue() != "Zebra" {
		t.Error("Expected value at the head of queue be Zebra")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("Cats")
	queue.Enqueue(7)
	queue.Enqueue("Dogs")

	value1 := queue.Dequeue()
	if value1 != "Cats" {
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
