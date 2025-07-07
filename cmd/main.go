package main

import (
	"fmt"
	"time"

	"github.com/emaforlin/qit/internal/queue"
)

func main() {
	messageQueue := queue.NewQueue("exampleQueue", queue.DefaultMaxQueueCapacity, queue.DefaultWokerCount)
	
	// Enqueue messages
	fmt.Println("Enqueueing messages...")
	for i := 0; i < 10; i++ {
		err := messageQueue.Enqueue(queue.Message{
			ID:      fmt.Sprintf("msg-%d", i),
			Payload: fmt.Sprintf("This is a test message with id %d", i),
		})
		if err != nil {
			fmt.Printf("Failed to enqueue message %d: %v\n", i, err)
		} else {
			fmt.Printf("Enqueued message: msg-%d\n", i)
		}
	}
	
	// Give workers some time to process messages
	fmt.Println("Waiting for workers to process messages...")
	time.Sleep(15 * time.Second) // Allow time for all messages to be processed
	
	// Close the queue and wait for workers to finish
	messageQueue.Close()
	messageQueue.Wait()
	
	fmt.Println("All messages processed. Exiting.")
}
