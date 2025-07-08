package broker

import (
	"sync"

	"github.com/emaforlin/qit/internal/queue"
)

type QueueBroker struct {
	mu     sync.RWMutex
	queues map[string]*queue.Queue
}

func NewQueueBroker() *QueueBroker {
	return &QueueBroker{
		queues: make(map[string]*queue.Queue),
	}
}

func (b *QueueBroker) GetQueue(name string) *queue.Queue {
	b.mu.RLock()
	q, exists := b.queues[name]
	b.mu.RUnlock()

	if exists {
		return q
	}
	return nil
}

func (b *QueueBroker) CreateQueue(name string, cap, workers int) *queue.Queue {
	b.mu.Lock()
	defer b.mu.Unlock()

	if q, exists := b.queues[name]; exists {
		return q
	}

	q := queue.NewQueue(name, cap, workers)
	b.queues[name] = q
	return q
}
