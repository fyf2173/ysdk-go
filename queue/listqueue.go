package queue

import (
	"container/list"
	"sync"
)

type LoopQueue struct {
	l *list.List
	m sync.Mutex
}

func NewLoopQueue() *LoopQueue {
	return &LoopQueue{
		l: list.New(),
	}
}

func (lq *LoopQueue) Enqueue(v ...interface{}) {
	if v == nil {
		return
	}
	lq.m.Lock()
	defer lq.m.Unlock()
	for _, e := range v {
		lq.l.PushBack(e)
	}
}

func (lq *LoopQueue) Dequeue() interface{} {
	lq.m.Lock()
	defer lq.m.Unlock()
	for e := lq.l.Front(); e != nil; {
		return lq.l.Remove(e)
	}
	return nil
}

func (lq *LoopQueue) Remove(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	lq.m.Lock()
	defer lq.m.Unlock()
	for e := lq.l.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			return lq.l.Remove(e)
		}
	}
	return nil
}

func (lq *LoopQueue) Len() int {
	lq.m.Lock()
	defer lq.m.Unlock()
	return lq.l.Len()
}
