package queue

type Element struct {
	data interface{}
	next *Element
}

type Queue struct {
	Head *Element
	Tail *Element
	Size int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (this *Queue) EnQueue(data interface{}) {
	e := new(Element)
	e.data = data
	if this.Size == 0 {
		this.Head = e
		this.Tail = e
	} else {
		this.Tail.next = e
	}
	this.Size++
	return
}

func (this *Queue) DeQueue() (data interface{}) {
	if this.Size == 0 {
		return
	}
	data = this.Head.data
	if this.Size == 1 {
		this.Head = nil
		this.Tail = nil
	} else {
		this.Head = this.Head.next
	}
	this.Size--
	return
}
