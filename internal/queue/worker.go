package queue

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id     int
	buffer <-chan Message
	wg     *sync.WaitGroup
}

// Start launches a goroutine that processes messages from the buffer channel.
func (w *Worker) Start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for msg := range w.buffer {
			fmt.Printf("[Worker %d] Processing message: %s\n", w.id, msg.ID)
			time.Sleep(1 * time.Second) // Simulate processing time
			fmt.Printf("[Worker %d] Finished processing message: %s\n", w.id, msg.ID)
		}
	}()
}

func NewWorker(id int, buffer <-chan Message, wg *sync.WaitGroup) *Worker {
	return &Worker{
		id:     id,
		buffer: buffer,
		wg:     wg,
	}
}
