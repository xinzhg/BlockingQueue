package BlockingQueue

import (
	"fmt"
	"testing"
	"time"
)

func TestNewBlockingQueue(t *testing.T) {
	bq := NewBlockingQueue(10)
	for i := 0; i < 10; i++ {
		bq.Enqueue(i)
	}
	go func() {
		time.Sleep(5 * time.Second)
		val, err := bq.Dequeue()
		fmt.Println("print val:", val, "err", err)
	}()
	bq.Enqueue(11)
	for i := 0; i < 11; i++ {
		func() {
			val, err := bq.Dequeue()
			fmt.Println("print val:", val, "err", err)
		}()
	}
}
