package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.front
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	n := ListItem{Value: v, Next: l.front, Prev: nil}
	if l.len != 0 {
		l.front.Prev = &n
	}
	l.front = &n
	if l.len == 0 {
		l.back = &n
	}
	l.len++
	return &n
}

func (l *list) PushBack(v interface{}) *ListItem {
	n := ListItem{Value: v, Next: nil, Prev: l.back}
	if l.len != 0 {
		l.back.Next = &n
	}
	l.back = &n
	if l.len == 0 {
		l.front = &n
	}
	l.len++
	return &n
}

func (l *list) Remove(i *ListItem) {
	if l.len <= 1 {
		l.front = nil
		l.back = nil
		l.len = 0
		return
	}
	// разбор краевых случаев
	switch i {
	case l.front:
		i.Next.Prev = nil
		l.front = i.Next
	case l.back:
		i.Prev.Next = nil
		l.back = i.Prev
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	i.Next = nil
	i.Prev = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l == nil || i == nil || l.len == 0 || i == l.front {
		return
	}
	if i.Prev != nil {
		// если не первый
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		// если не последний
		i.Next.Prev = i.Prev
	}
	if l.back == i {
		// если последний
		l.back = i.Prev
	}
	i.Prev = nil
	i.Next = l.front
	if l.front != nil {
		l.front.Prev = i
	}
	l.front = i
}

func NewList() List {
	return new(list)
}
