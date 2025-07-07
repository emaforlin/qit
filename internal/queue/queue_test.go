package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	mq := NewMessageQueue()
	if mq == nil {
		t.Error("Expected a new MessageQueue instance, got nil")
	}
}

func TestEnqueue(t *testing.T) {
	mq := NewMessageQueue()
	msg := Message{ID: "1", Payload: "test message"}

	err := mq.Enqueue(msg)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mq.queue) != 1 {
		t.Errorf("Expected queue length to be 1, got %d", len(mq.queue))
	}
}

func TestDequeue(t *testing.T) {
	mq := NewMessageQueue()
	msg := Message{ID: "1", Payload: "test message"}

	err := mq.Enqueue(msg)
	if err != nil {
		t.Fatalf("Expected no error on enqueue, got %v", err)
	}

	dequeuedMsg, err := mq.Dequeue()
	if err != nil {
		t.Errorf("Expected no error on dequeue, got %v", err)
	}

	if dequeuedMsg.ID != msg.ID || dequeuedMsg.Payload != msg.Payload {
		t.Errorf("Expected dequeued message to match enqueued message, got %v", dequeuedMsg)
	}

	if len(mq.queue) != 0 {
		t.Errorf("Expected queue length to be 0 after dequeue, got %d", len(mq.queue))
	}
}
