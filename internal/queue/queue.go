package queue

import (
	"errors"
	"sync"
)

const (
	DefaultQueueCapacity    = 100 // default initial capacity of the message queue
	DefaultMaxQueueCapacity = 200 // maximum capacity of the message queue

	// Error messages
	EmptyQueueError = "queue is empty" // error message when trying to dequeue from an empty queue
	FullQueueError  = "queue is full"  // error message when trying to enqueue to a full queue
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
		mu:    sync.Mutex{},
		queue: make([]Message, 0, DefaultQueueCapacity),
	}
}

// Enqueue adds a message to the queue.
func (mq *MessageQueue) Enqueue(msg Message) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if len(mq.queue) >= DefaultMaxQueueCapacity {
		return errors.New(FullQueueError)
	}

	mq.queue = append(mq.queue, msg)
	return nil
}

// Dequeue removes and returns the first (oldest) message from the queue.
func (mq *MessageQueue) Dequeue() (Message, error) {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	if len(mq.queue) == 0 {
		return Message{}, errors.New(EmptyQueueError)
	}
	msg := mq.queue[0]
	mq.queue = mq.queue[1:]
	return msg, nil
}
