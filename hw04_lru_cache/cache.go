package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if i, ok := l.items[key]; ok {
		i.Value = value
		l.queue.MoveToFront(i)
		return true
	}
	l.items[key] = l.queue.PushFront(value)
	if l.queue.Len() > l.capacity {
		tail := l.queue.Back()
		if tail != nil {
			l.queue.Remove(tail)
			for k, v := range l.items {
				if v == tail {
					delete(l.items, k)
					break
				}
			}
		}
	}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if i, ok := l.items[key]; ok {
		l.queue.MoveToFront(i)
		return i.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.items = make(map[Key]*ListItem)
	for l.queue.Len() > 0 {
		l.queue.Remove(l.queue.Front())
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
