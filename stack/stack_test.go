package stack

import "testing"

func TestStack_Empty(t *testing.T) {
	s := NewStack()
	if !s.Empty() {
		t.Fatalf("NewStack isn't empty")
	}

	s.Push(1)
	s.Pop()
	if !s.Empty() {
		t.Fatalf("Stack push 1 and pop isn't empty")
	}
}

func TestStack_Peek(t *testing.T) {
	s := NewStack()
	s.Push(1)
	if s.Peek() != 1 {
		t.Fatalf("Stack top element isn't 1")
	}

	s.Push(2)
	if s.Peek() != 2 {
		t.Fatalf("Stack top element isn't 2")
	}
}

func TestStack_Size(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Fatalf("NewStack size isn't 0")
	}

	s.Push(1)
	if s.Size() != 1 {
		t.Fatalf("Stack size isn't 1")
	}

	s.Pop()
	if s.Size() != 0 {
		t.Fatalf("Stack size isn't 0")
	}
}

func TestStack_Pop(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("Stack Pop crashed")
		}
	}()
	s := NewStack()
	s.Push(1)
	s.Pop()
	s.Pop()
}
