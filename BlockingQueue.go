package BlockingQueue

import (
	"errors"
	"sync"
)

/*
	Blocking queue for multi-goroutines
*/
type BlockingQueue struct {
	size int
	cond *sync.Cond
	data []int
}

func NewBlockingQueue(size int) *BlockingQueue {
	return &BlockingQueue{size, sync.NewCond(&sync.Mutex{}), make([]int, 0, size)}
}

/*
	It is a blocking method
*/
func (b *BlockingQueue) Enqueue(val int) {
	if func() bool {
		b.cond.L.Lock()
		defer b.cond.L.Unlock()
		if len(b.data) < b.size {
			b.data = append(b.data, val)
			return true
		}
		return false
	}() {
		return
	}
	b.cond.L.Lock()
	b.cond.Wait()
	defer b.cond.L.Unlock()
	b.data = append(b.data, val)
}

func (b *BlockingQueue) Dequeue() (int, error) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	if len(b.data) == 0 {
		return 0, errors.New("no data")
	}
	front := b.data[0]
	b.data = b.data[1:]
	b.cond.Signal()
	return front, nil
}
