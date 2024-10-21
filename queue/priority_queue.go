package queue

import "container/heap"

type PriorityQueue struct {
	heap *MaxHeap
}

func NewPriorityQueue() *PriorityQueue {
	heapStructure := make(MaxHeap, 0)
	return &PriorityQueue{
		heap: &heapStructure,
	}
}

func NewPriorityQueueWithData(inputs ...HeapData) *PriorityQueue {
	heapStructure := MaxHeap(inputs)
	heap.Init(&heapStructure)
	return &PriorityQueue{
		heap: &heapStructure,
	}
}

func (r *PriorityQueue) Push(value HeapData) {
	heap.Push(r.heap, value)
}

func (r *PriorityQueue) Pop() *HeapData {
	poppedValue := heap.Pop(r.heap)
	output := poppedValue.(HeapData)
	return &output
}

func (r *PriorityQueue) Size() int {
	return len(*r.heap)
}

type HeapData struct {
	value    interface{}
	priority int
}

type MaxHeap []HeapData

func (r *MaxHeap) Len() int {
	return len(*r)
}

func (r *MaxHeap) Less(i, j int) bool {
	heapList := *r
	return heapList[i].priority > heapList[j].priority
}

func (r *MaxHeap) Swap(i, j int) {
	heapList := *r
	heapList[i], heapList[j] = heapList[j], heapList[i]
}

func (r *MaxHeap) Push(value interface{}) {
	data := value.(HeapData)
	*r = append(*r, data)
}

func (r *MaxHeap) Pop() interface{} {
	heapList := *r
	heapListLength := len(heapList)
	if heapListLength == 0 {
		return nil
	}

	data := heapList[heapListLength-1]
	*r = heapList[:heapListLength-1]
	return data
}
