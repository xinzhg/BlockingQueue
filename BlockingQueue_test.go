package BlockingQueue

import (
	"reflect"
	"testing"
)

func TestBlockingQueue_Dequeue(t *testing.T) {
	type fields struct {
		size int
		cond *sync.Cond
		data []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlockingQueue{
				size: tt.fields.size,
				cond: tt.fields.cond,
				data: tt.fields.data,
			}
			got, err := b.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockingQueue_Enqueue(t *testing.T) {
	type fields struct {
		size int
		cond *sync.Cond
		data []int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlockingQueue{
				size: tt.fields.size,
				cond: tt.fields.cond,
				data: tt.fields.data,
			}
			b.Enqueue(tt.args.val)
		})
	}
}

func TestNewBlockingQueue(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *BlockingQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlockingQueue(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlockingQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
