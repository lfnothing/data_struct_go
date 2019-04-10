package stack

type Stack struct {
	data []interface{}
	dataLength int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0, 100),
		dataLength:0,
	}
}

func (this *Stack) Push(value interface{}) {
	this.data = append(this.data, value)
	this.dataLength++
}

func (this *Stack) Pop() {
	if this.dataLength != 0 {
		this.data = this.data[:this.dataLength-1]
		this.dataLength--
	}
}

func (this *Stack) Peek() (value interface{}) {
	if this.dataLength == 0 {
		return
	}
	return this.data[this.dataLength - 1]
}

func (this *Stack) Empty() bool {
	return this.dataLength == 0
}

func (this *Stack) Size() int {
	return this.dataLength
}