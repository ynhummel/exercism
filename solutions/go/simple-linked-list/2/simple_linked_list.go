package linkedlist

import (
	"errors"
)

// Define the List and Element types here.
type Element struct {
	val  int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(elements []int) *List {
	list := &List{}
	for _, ele := range elements {
		list.Push(ele)
	}

	return list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	newElement := &Element{
		val:  element,
		next: l.head,
	}
	l.head = newElement
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size < 1 {
		return 0, errors.New("Can not pop from empty list")
	}

	val := l.head.val
	l.head = l.head.next
	l.size--
	return val, nil
}

func (l *List) Array() []int {
	result := make([]int, l.size)
	currEl := l.head
	for i := l.size - 1; i >= 0; i-- {
		result[i] = currEl.val
		currEl = currEl.next
	}

	return result
}

func (l *List) Reverse() *List {
	revList := &List{}
	currEl := l.head
	for currEl != nil {
		revList.Push(currEl.val)
		currEl = currEl.next
	}

	return revList
}
