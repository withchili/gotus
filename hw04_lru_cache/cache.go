package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheElement struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	newElement := cacheElement{
		key:   key,
		value: value,
	}
	if i, ok := l.items[key]; ok {
		i.Value = newElement
		l.queue.MoveToFront(i)
		return true
	}
	l.items[key] = l.queue.PushFront(newElement)
	if l.queue.Len() > l.capacity {
		tail := l.queue.Back()
		if tail != nil {
			l.queue.Remove(tail)
			delete(l.items, tail.Value.(cacheElement).key)
		}
	}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if i, ok := l.items[key]; ok {
		l.queue.MoveToFront(i)
		return i.Value.(cacheElement).value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.items = make(map[Key]*ListItem)
	for l.queue.Len() > 0 {
		l.queue.Remove(l.queue.Front())
	}
}
