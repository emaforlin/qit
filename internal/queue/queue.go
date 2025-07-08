package queue

import (
	"errors"
	"sync"
)

const (
	DefaultMaxQueueCapacity = 200 // maximum capacity of the message queue
	DefaultWokerCount       = 5   // default number of workers to process messages in the queue

	// Error messages
	EmptyQueueError = "queue is empty"
	FullQueueError  = "queue is full"
)

type Queue struct {
	wg       sync.WaitGroup
	name     string
	buffer   chan Message
	workers  []*Worker
	capacity int
}

func NewQueue(name string, capacity, numWorkers int) *Queue {
	q := &Queue{
		name:     name,
		buffer:   make(chan Message, capacity),
		workers:  make([]*Worker, 0, numWorkers),
		capacity: capacity,
	}

	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i, q.buffer, &q.wg)
		q.workers = append(q.workers, worker)
		worker.Start()
	}
	return q
}

// Enqueue adds a message to the queue.
func (mq *Queue) Enqueue(msg Message) error {
	select {
	case mq.buffer <- msg:
		return nil
	default:
		return errors.New(FullQueueError)
	}
}

// Close closes the queue and stops accepting new messages
func (mq *Queue) Close() {
	close(mq.buffer)
}

// Wait waits for all workers to finish processing their current messages
func (mq *Queue) Wait() {
	mq.wg.Wait()
}
