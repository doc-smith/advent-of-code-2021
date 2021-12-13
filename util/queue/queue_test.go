package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if !q.Empty() {
		t.Error("new queue must be empty")
	}
}
