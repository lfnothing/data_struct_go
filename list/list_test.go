package list

import (
	"fmt"
	"strconv"
	"testing"
)

type ListInt struct {
	i int
}

func NewListInt() *ListInt {
	return new(ListInt)
}

func (this *ListInt) I(i int) *ListInt {
	this.i = i
	return this
}

func (this *ListInt) destory(i interface{}) {
	fmt.Printf("destoryed %d\n", i.(ListInt).i)
}

func (this *ListInt) equal(i1 interface{}, i2 interface{}) bool {
	return i1.(ListInt).i == i2.(ListInt).i
}

func (this ListInt) String() string {
	return strconv.Itoa(this.i) + " "
}

func TestNewList(t *testing.T) {
	l := NewList(NewListInt())
	if !l.equal(NewListInt().I(1), NewListInt().I(1)) {
		t.Fatalf("List equal error: 1 != 1\n")
	}
	l.destory(1)
	l.destory(2)
}

func TestList_Insert(t *testing.T) {
	l1 := NewList(NewListInt())
	listInt := []ListInt{{1}, {2}, {3}, {4}, {5}}

	for i := 0; i < len(listInt); i++ {
		l1.Insert(l1.tail, listInt[i])
	}
	for e := l1.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l2 := NewList(NewListInt())
	for i := 0; i < len(listInt); i++ {
		l2.Insert(nil, listInt[i])
	}
	for e := l2.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()
}

func TestList_Delete(t *testing.T) {
	l := NewList(NewListInt())
	listInt := []ListInt{{1}, {2}, {3}, {4}, {5}}

	for i := 0; i < len(listInt); i++ {
		l.Insert(l.tail, listInt[i])
	}
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(listInt[0])
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(listInt[1])
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(listInt[2])
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(listInt[3])
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(listInt[4])
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(listInt[0])
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()
}

func TestList_Update(t *testing.T) {
	update := ListInt{i: 6}
	l := NewList(NewListInt())
	listInt := []ListInt{{1}, {2}, {3}, {4}, {5}}

	for i := 0; i < len(listInt); i++ {
		l.Insert(l.tail, listInt[i])
	}
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Update(listInt[0], update)
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Update(listInt[1], update)
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Update(listInt[2], update)
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Update(listInt[3], update)
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	l.Update(listInt[4], update)
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()
}

func TestList_Find(t *testing.T) {
	l := NewList(NewListInt())
	listInt := []ListInt{{1}, {2}, {3}, {4}, {5}}

	for i := 0; i < len(listInt); i++ {
		l.Insert(l.tail, listInt[i])
	}
	for e := l.head; e != nil; e = e.next {
		fmt.Printf(e.data.(ListInt).String())
	}
	fmt.Println()

	e := l.Find(listInt[0])
	fmt.Println(e.data.(ListInt).String())

	e = l.Find(listInt[1])
	fmt.Println(e.data.(ListInt).String())

	e = l.Find(listInt[2])
	fmt.Println(e.data.(ListInt).String())

	e = l.Find(listInt[3])
	fmt.Println(e.data.(ListInt).String())

	e = l.Find(listInt[4])
	fmt.Println(e.data.(ListInt).String())
}
