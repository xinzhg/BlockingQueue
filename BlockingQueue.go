package BlockingQueue

import "sync"

type BlockingQueue struct {
	size int
	cond sync.Cond
}
