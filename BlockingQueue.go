package BlockingQueue

import "sync"

type BlockingQueue struct {
	size int
	cond *sync.Cond
}

func NewBlockingQueue(size int) *BlockingQueue {
	return &BlockingQueue{size, sync.NewCond(&sync.Mutex{})}
}
