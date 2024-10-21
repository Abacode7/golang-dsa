package queue

import (
	"fmt"
	"math"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	priorityQueue := NewPriorityQueueWithData(
		HeapData{"Aston Villa", 7},
		HeapData{"Chelsea", 2},
		HeapData{"Arsenal", 4},
		HeapData{"Liverpool", 3},
		HeapData{"Manchester United", 1},
	)

	priorityQueue.Push(HeapData{"Manchester City", 6})

	priority := math.MaxInt
	for priorityQueue.Size() != 0 {
		poppedData := priorityQueue.Pop()

		if poppedData.priority > priority {
			t.Error("Priority Queue does not process values in the order of maximum priority")
		}
		priority = poppedData.priority

		fmt.Println(poppedData.priority, poppedData.value)
	}
}
