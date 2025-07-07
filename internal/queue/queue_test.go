package queue

import (
	"testing"
)

const (
	TestingMaxQueueCap = 20
	TestingWorkerCount = 2
)

func TestNewQueue(t *testing.T) {
	mq := NewQueue("testQueue", TestingMaxQueueCap, TestingWorkerCount)
	if mq == nil {
		t.Error("Expected a new Queue instance, got nil")
	}
}

func TestEnqueue(t *testing.T) {
	mq := NewQueue("testQueue", TestingMaxQueueCap, TestingWorkerCount)
	msg := Message{ID: "1", Payload: "test message"}

	err := mq.Enqueue(msg)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mq.buffer) != 1 {
		t.Errorf("Expected queue length to be 1, got %d", len(mq.buffer))
	}
}
