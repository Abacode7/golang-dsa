package stack

import "testing"

func TestNewStack(t *testing.T) {
	t.Run("Test NewStack returns valid stack", func(t *testing.T) {

	})

	t.Run("Test NewStack has valid underlying data", func(t *testing.T) {

	})
}

func TestStack_Push(t *testing.T) {

	t.Run("Test Push performs as expected", func(t *testing.T) {
		stack := NewStack()
		stack.Push("a")
		stack.Push("b")

		if stack.Peek() != "b" {
			t.Error("Expected value not returned from stack, got this instead: ", stack.Peek())
		}
	})

	t.Run("Test Push performs as expected on an empty stack", func(t *testing.T) {
		stack := NewStack()
		stack.Push("z")
		if stack.Peek() != "z" {
			t.Error("Expected value not return from stack, got this instead: ", stack.Peek())
		}
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("Test Pop pops expected value", func(t *testing.T) {
		stack := NewStack()
		stack.Push(6)

		poppedValue := stack.Pop()
		if poppedValue != 6 {
			t.Error("Expected value not return from stack, got this instead: ", poppedValue)
		}
	})

	t.Run("Test Pop returns nil when stack is empty", func(t *testing.T) {
		stack := NewStack()
		if stack.Pop() != nil {
			t.Error("Expected popped value from empty stack be nil")
		}
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("Test Peek returns expected value", func(t *testing.T) {
		stack := NewStack()
		stack.Push("Pawpaw")
		stack.Push("Mango")
		stack.Push("Avocado")

		if stack.Peek() != "Avocado" {
			t.Error("Expected value not returned, got this instead: ", stack.Peek())
		}
	})

	t.Run("Test Pekk returns nil for empty stack", func(t *testing.T) {
		stack := NewStack()

		if stack.Peek() != nil {
			t.Error("Expected nil value upon peeking empty stack")
		}
	})
}
