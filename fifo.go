package fifo

import "sync/atomic"

type FIFO[T any] struct {
	queue []T
	pos   atomic.Uint32
}

func NewFIFO[T any](size int) *FIFO[T] {
	fifo := &FIFO[T]{queue: make([]T, size)}
	fifo.pos.Store(^uint32(0))
	return fifo
}

func (fifo *FIFO[T]) Push(n T) {
	pos := fifo.pos.Add(1) % uint32(len(fifo.queue))
	fifo.queue[int(pos)] = n
}

func (fifo *FIFO[T]) Each(cb func(T)) {
	for _, v := range fifo.queue {
		cb(v)
	}
}
