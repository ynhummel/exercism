package linkedlist

import (
	"errors"
)

// Define the List and Element types here.
type Element struct {
	val  int
	Next *Element
}

type List struct {
	head *Element
}

func New(elements []int) *List {
	list := &List{head: nil}
	if len(elements) < 1 {
		return list
	}

	for _, el := range elements {
		oldHead := list.head
		list.head = &Element{
			val:  el,
			Next: oldHead,
		}
	}

	return list
}

func (l *List) Size() int {
	if l.head == nil {
		return 0
	}

	size := 1
	next := l.head.Next
	for next != nil {
		size++
		next = next.Next
	}

	return size
}

func (l *List) Push(element int) {
	newElement := &Element{
		val:  element,
		Next: l.head,
	}

	l.head = newElement
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("Can not pop from empty list")
	}
	val := l.head.val
	l.head = l.head.Next

	return val, nil
}

func (l *List) Array() []int {
	llen := l.Size()
	result := make([]int, llen)
	if llen == 0 {
		return result
	}

	ele := l.head
	for i := llen - 1; i >= 0; i-- {
		result[i] = ele.val
		ele = ele.Next
	}

	return result
}

func (l *List) Reverse() *List {
	if l.head == nil || l.Size() == 1 {
		return l
	}

	revList := &List{head: nil}

	list := l.head
	for list != nil {
		revList.Push(list.val)
		list = list.Next
	}

	return revList
}
