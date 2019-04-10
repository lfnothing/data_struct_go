package circular_queue

import (
	"testing"
)

func TestCircularQueue_Empty(t *testing.T) {
	q := NewCircularQueue(10)
	if !q.Empty() {
		t.Fatalf("New circular queue isn't empty")
	}
	q.EnQueue(1)
	q.DeQueue()
	if !q.Empty() {
		t.Fatalf("Circual queue isn't empty")
	}
}

func TestCircularQueue_Full(t *testing.T) {
	q := NewCircularQueue(2)
	q.EnQueue(1)
	q.EnQueue(2)
	if !q.Full() {
		t.Fatalf("Circual queue isn't full")
	}
}

func TestCircularQueue_Front(t *testing.T) {
	q := NewCircularQueue(2)
	q.EnQueue(1)
	if q.Front() != 1 {
		t.Fatalf("Circular queue front is not 1")
	}
	q.DeQueue()
	q.EnQueue(2)
	if q.Front() != 2 {
		t.Fatalf("Circular queue front is not 2")
	}
	q.EnQueue(3)
	q.EnQueue(4)
	q.DeQueue()
	q.DeQueue()
	q.EnQueue(5)
	if q.Front() != 5 {
		t.Fatalf("Circular queue front is not 5")
	}
}

func TestCircularQueue_Rear(t *testing.T) {
	q := NewCircularQueue(2)
	q.EnQueue(1)
	if q.Rear() != 1 {
		t.Fatalf("Circular queue tail is not 1")
	}
	q.DeQueue()
	q.EnQueue(2)
	if q.Rear() != 2 {
		t.Fatalf("Circular queue tail is not 2")
	}
	q.EnQueue(3)
	q.EnQueue(4)
	q.DeQueue()
	q.DeQueue()
	q.EnQueue(5)
	if q.Rear() != 5 {
		t.Fatalf("Circular queue tail is not 5")
	}
}

func TestCircularQueue_EnQueue(t *testing.T) {
	q := NewCircularQueue(2)
	q.EnQueue(1)
	q.EnQueue(2)
	if q.EnQueue(3) {
		t.Fatalf("Circular queue is full can't enqueue")
	}
}

func TestCircularQueue_DeQueue(t *testing.T) {
	q := NewCircularQueue(2)
	if q.DeQueue() {
		t.Fatalf("Circular queue is empty can't dequeue")
	}
	q.EnQueue(1)
	q.DeQueue()
	if q.DeQueue() {
		t.Fatalf("Circular queue is empty can't dequeue")
	}
}
