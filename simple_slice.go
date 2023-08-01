package simpleslice

import "sync/atomic"

type SimpleSlice[T any] struct {
	queue []T
	pos   atomic.Uint32
}

func NewSimpleSlice[T any](size int) *SimpleSlice[T] {
	fifo := &SimpleSlice[T]{queue: make([]T, size)}
	fifo.pos.Store(^uint32(0))
	return fifo
}

func (slice *SimpleSlice[T]) Push(n T) {
	pos := slice.pos.Add(1) % uint32(len(slice.queue))
	slice.queue[int(pos)] = n
}

func (slice *SimpleSlice[T]) Each(cb func(T)) {
	for _, v := range slice.queue {
		cb(v)
	}
}
