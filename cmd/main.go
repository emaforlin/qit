package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	handlers "github.com/emaforlin/qit/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	srv := setupServer()

	// Start server in goroutine
	go startServer(srv)

	// Wait for shutdown signal and gracefully close
	gracefulShutdown(srv)

	// messageQueue := queue.NewQueue("exampleQueue", queue.DefaultMaxQueueCapacity, queue.DefaultWokerCount)

	// // Enqueue messages
	// fmt.Println("Enqueueing messages...")
	// for i := 0; i < 10; i++ {
	// 	err := messageQueue.Enqueue(queue.Message{
	// 		ID:      fmt.Sprintf("msg-%d", i),
	// 		Payload: fmt.Sprintf("This is a test message with id %d", i),
	// 	})
	// 	if err != nil {
	// 		fmt.Printf("Failed to enqueue message %d: %v\n", i, err)
	// 	} else {
	// 		fmt.Printf("Enqueued message: msg-%d\n", i)
	// 	}
	// }

	// // Give workers some time to process messages
	// fmt.Println("Waiting for workers to process messages...")
	// time.Sleep(15 * time.Second) // Allow time for all messages to be processed

	// // Close the queue and wait for workers to finish
	// messageQueue.Close()
	// messageQueue.Wait()

	// fmt.Println("All messages processed. Exiting.")
}

// setupServer creates and configures the HTTP server
func setupServer() *http.Server {
	router := gin.Default()
	router.POST("/messages", handlers.PostMessage)

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}

// startServer starts the HTTP server
func startServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

// gracefulShutdown waits for shutdown signal and gracefully closes the server
func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 seconds.")
	default:
		log.Println("Server exiting gracefully")
	}
}
