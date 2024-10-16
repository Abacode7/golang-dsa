package linkedlist

import (
	"testing"
)

func TestSinglyLinkedList_GetHead(t *testing.T) {
	linkedList := NewSinglyLinkedList()

	firstValue := "1"
	secondValue := "2"

	linkedList.InsertAt(0, NewSNode(firstValue, nil))
	linkedList.InsertAt(1, NewSNode(secondValue, nil))

	t.Run("Test GetHead on populated linked list should return expected value", func(t *testing.T) {
		head := linkedList.GetHead()
		if head.value != firstValue {
			t.Error("Expected head value to be: ", firstValue)
		}
	})

	t.Run("Test GetHead on empty linked list should return nil value", func(t *testing.T) {
		newLinkedList := NewSinglyLinkedList()
		newHead := newLinkedList.GetHead()
		if newHead != nil {
			t.Error("Expected head for empty linked list be nil, but got: ", newHead.value)
		}
	})
}

func TestSinglyLinkedList_GetAt(t *testing.T) {
	linkedList := NewSinglyLinkedList()

	firstValue := "1"
	secondValue := "2"

	linkedList.InsertAt(0, NewSNode(firstValue, nil))
	linkedList.InsertAt(1, NewSNode(secondValue, nil))

	t.Run("Test GetAt first index should return expected value", func(t *testing.T) {
		if linkedList.GetAt(0).value != firstValue {
			t.Error("Expected linked list first value be: ", firstValue)
		}
	})

	t.Run("Test GetAt second index should return expected value", func(t *testing.T) {
		if linkedList.GetAt(1).value != secondValue {
			t.Error("Expected linked list first value be: ", secondValue)
		}
	})

	t.Run("Test GetAt index out of bound should return nil", func(t *testing.T) {
		if linkedList.GetAt(10) != nil {
			t.Error("Expected linked list index out of bounds be defaulted to last value: ", linkedList.GetAt(10).value)
		}
	})
}

func TestSinglyLinkedList_RemoveAt(t *testing.T) {
	linkedList := NewSinglyLinkedList()

	firstValue := "5"
	secondValue := "10"
	thirdValue := "15"

	linkedList.InsertAt(0, NewSNode(firstValue, nil))
	linkedList.InsertAt(1, NewSNode(secondValue, nil))
	linkedList.InsertAt(2, NewSNode(thirdValue, nil))

	t.Run("RemoveAt head returns expected value", func(t *testing.T) {
		removedValue := linkedList.RemoveAt(0)
		if removedValue.value != firstValue {
			t.Error("Expected value not removed at head, got this instead: ", removedValue.value)
		}
	})

	t.Run("RemoveAt index out of bounds returns nil value", func(t *testing.T) {
		removedValue := linkedList.RemoveAt(10)
		if removedValue != nil {
			t.Error("Expected value not removed at head should be nil")
		}
	})

	t.Run("RemoveAt valid index returns expected value", func(t *testing.T) {
		removedValue := linkedList.RemoveAt(1)
		if removedValue.value != thirdValue {
			t.Error("Expected value not removed at index, got this instead: ", removedValue.value)
		}
	})

	t.Run("RemoveAt head in empty linked list returns nil", func(t *testing.T) {
		linkedList = NewSinglyLinkedList()
		removedValue := linkedList.RemoveAt(0)
		if removedValue != nil {
			t.Error("Expected value removed at head should be nil")
		}
	})
}
