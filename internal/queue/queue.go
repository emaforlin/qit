package queue

import "sync"

const (
	DefaultQueueSize = 100
)

type Message struct {
	ID      string
	Payload any
}

type MessageQueue struct {
	mu    sync.Mutex
	queue []Message
}

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		queue: make([]Message, 0, DefaultQueueSize),
	}
}
