package circular_queue

type CircularQueue struct {
	data []interface{}
	head int
	tail int
}

func NewCircularQueue(k int) *CircularQueue {
	return &CircularQueue{
		data: make([]interface{}, k, k),
		head: -1,
		tail: -1,
	}
}

func (this *CircularQueue) EnQueue(value interface{}) bool {
	if this.Full() {
		return false
	}
	if this.Empty() {
		this.head = 0
	}
	this.tail = (this.tail + 1) % cap(this.data)
	this.data[this.tail] = value
	return true
}

func (this *CircularQueue) DeQueue() bool {
	if this.Empty() {
		return false
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}
	this.head = (this.head + 1) % cap(this.data)
	return true
}

func (this *CircularQueue) Front() (v interface{}) {
	if this.Empty() {
		return
	}
	return this.data[this.head]
}

func (this *CircularQueue) Rear() (v interface{}) {
	if this.Empty() {
		return
	}
	return this.data[this.tail]
}

func (this *CircularQueue) Empty() bool {
	if this.head == -1 {
		return true
	}
	return false
}

func (this *CircularQueue) Full() bool {
	return (this.tail+1)%cap(this.data) == this.head
}
