package heap

import "container/heap"

type Person struct {
	Name string
	Age  int
}

// PersonList implements sort.Interface based on the Age field.
type PersonList []Person

func NewPersonList(name ...Person) *PersonList {
	personList := PersonList(name)
	heap.Init(&personList)
	return &personList
}

func (a PersonList) Len() int { return len(a) }

func (a PersonList) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a PersonList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a *PersonList) Push(x interface{}) {
	*a = append(*a, x.(Person))
}

func (a *PersonList) Pop() interface{} {
	personList := *a
	if len(personList) == 0 {
		return nil
	}

	n := len(personList)
	poppedPerson := personList[n-1]
	*a = personList[0 : n-1]
	return poppedPerson
}

func (a *PersonList) PushHeap(x interface{}) {
	heap.Push(a, x)
}

func (a *PersonList) PopHeap() interface{} {
	return heap.Pop(a)
}
