package dlist

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

func (this *ListInt) Destory(i interface{}) {
	if v, ok := i.(ListInt); ok {
		fmt.Printf("destoryed %d\n", v.i)
	} else if v, ok := i.(ListInt); ok {
		fmt.Printf("destoryed %d\n", v.i)
	}
}

func (this *ListInt) Equal(i1 interface{}, i2 interface{}) bool {
	if v1, ok := i1.(ListInt); ok {
		if v2, ok := i2.(ListInt); ok {
			return v1.i == v2.i
		}

		if v2, ok := i2.(*ListInt); ok {
			return v1.i == v2.i
		}
	} else if v1, ok := i1.(*ListInt); ok {
		if v2, ok := i2.(ListInt); ok {
			return v1.i == v2.i
		}

		if v2, ok := i2.(*ListInt); ok {
			return v1.i == v2.i
		}
	}
	return false
}

func (this ListInt) String() string {
	return strconv.Itoa(this.i) + " "
}

func TestNewList(t *testing.T) {
	l := NewList(NewListInt())
	if !l.Equal(NewListInt().I(1), NewListInt().I(1)) {
		t.Fatalf("List equal error: 1 != 1\n")
	}
}

func TestList_Insert(t *testing.T) {
	l1 := NewList(NewListInt())
	listInt := []ListInt{{1}, {2}, {3}, {4}, {5}}

	for i := 0; i < len(listInt); i++ {
		l1.Insert(l1.Tail, listInt[i])
	}
	for e := l1.Head; e != nil; e = e.Next {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
	for e := l1.Tail; e != nil; e = e.Pre {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()

	l2 := NewList(NewListInt())
	for i := 0; i < len(listInt); i++ {
		l2.Insert(nil, listInt[i])
	}
	for e := l2.Head; e != nil; e = e.Next {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
	for e := l2.Tail; e != nil; e = e.Pre {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
}

func TestList_Delete(t *testing.T) {
	l := NewList(NewListInt())
	listInt := []ListInt{{1}, {2}, {3}, {4}, {5}}

	for i := 0; i < len(listInt); i++ {
		l.Insert(l.Tail, listInt[i])
	}
	for e := l.Head; e != nil; e = e.Next {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
	for e := l.Tail; e != nil; e = e.Pre {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(l.Head)
	for e := l.Head; e != nil; e = e.Next {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
	for e := l.Tail; e != nil; e = e.Pre {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(l.Head)
	for e := l.Head; e != nil; e = e.Next {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
	for e := l.Tail; e != nil; e = e.Pre {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(l.Head)
	for e := l.Head; e != nil; e = e.Next {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
	for e := l.Tail; e != nil; e = e.Pre {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(l.Head)
	for e := l.Head; e != nil; e = e.Next {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
	for e := l.Tail; e != nil; e = e.Pre {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()

	l.Delete(l.Head)
	for e := l.Head; e != nil; e = e.Next {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
	for e := l.Tail; e != nil; e = e.Pre {
		fmt.Printf(e.Data.(ListInt).String())
	}
	fmt.Println()
}
