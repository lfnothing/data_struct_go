package list

type list interface {
	destory(interface{})
	equal(interface{}, interface{}) bool
}

type List struct {
	head *Element
	tail *Element
	size int
	list
}

type Element struct {
	data interface{}
	next *Element
}

func NewList(l list) *List {
	return &List{
		head: nil,
		tail: nil,
		size: 0,
		list: l,
	}
}

func (this *List) Insert(pre *Element, data interface{}) (e *Element) {
	e = new(Element)
	e.data = data

	if pre == nil {
		e.next = this.head
		this.head = e
		if this.tail == nil {
			this.tail = this.head
		}
	} else {
		e.next = pre.next
		pre.next = e
		if this.tail == pre {
			this.tail = e
		}
	}
	this.size++
	return e
}

func (this *List) Delete(data interface{}) {
	if this.size == 0 {
		return
	}

	var find bool
	if this.equal(this.head.data, data) {
		find = true
		this.head = this.head.next
		if this.head == nil {
			this.tail = nil
		}
	} else {
		for pre := this.head; pre != nil; pre = pre.next {
			if pre.next != nil && this.equal(pre.next.data, data) {
				find = true
				pre.next = pre.next.next
				if pre.next == nil {
					this.tail = pre
				}
			}
		}
	}

	if find {
		this.destory(data)
		this.size--
	}
	return
}

func (this *List) Update(old interface{}, new interface{}) {
	find := this.Find(old)
	if find != nil {
		find.data = new
	}
}

func (this *List) Find(d interface{}) (e *Element) {
	for e = this.head; e != nil; e = e.next {
		if this.equal(e.data, d) {
			break
		}
	}
	return
}
