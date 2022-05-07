package BlockingQueue

import (
	"errors"
	"sync"
)

type BlockingQueue struct {
	size int
	cond *sync.Cond
	data []int
}

func NewBlockingQueue(size int) *BlockingQueue {
	return &BlockingQueue{size, sync.NewCond(&sync.Mutex{}), make([]int, 0, size)}
}

func (b *BlockingQueue) Enqueue(val int) error {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	if len(b.data) == b.size {
		return errors.New("reach to the limit size")
	}
	b.data = append(b.data, val)
	return nil
}

func (b *BlockingQueue) Dequeue() (int, error) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	if len(b.data) == 0 {
		return 0, errors.New("no data")
	}
	front := b.data[0]
	b.data = b.data[1:]
	return front, nil
}
