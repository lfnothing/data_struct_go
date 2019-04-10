package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Fatalf("NewQueue() failed: return nil")
	}
	if q.Head != nil {
		t.Fatalf("New queue head != nil")
	}
	if q.Tail != nil {
		t.Fatalf("New queue tail != nil")
	}
	if q.Size != 0 {
		t.Fatalf("New queue size != 0")
	}
}

func TestQueue_EnDequeue(t *testing.T) {
	q := NewQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	if q.Size != 2 {
		t.Fatalf("Enqueue failed")
	}
	if q.DeQueue() != 1 && q.Size != 1 {
		t.Fatalf("Enqueue Dequeue failed")
	}
	if q.DeQueue() != 2 && q.Size != 0 {
		t.Fatalf("Enqueue Dequeue failed")
	}
	if q.DeQueue() != nil {
		t.Fatalf("Dequeue failed")
	}
}
