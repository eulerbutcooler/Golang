package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
 * INFO: Context Cancellation (Deadlines, Propagation, Cleanup)
 The context package is Go's standard way to scope requests and operations. It carries deadlines,
 cancellation signals, and request-scoped values across API boundaries and between goroutines.
 The Tree Structure: Contexts form a tree. You start with context.Background(). When you derive
 a new context (e.g., using context.WithCancel(parent)), you create a child node.
 Propagation: If a parent context is canceled, all of its children are instantly canceled.
 However, canceling a child does not cancel the parent.
 Deadlines: WithTimeout and WithDeadline are just auto-canceling contexts based on the clock.
 The Cleanup Rule (Crucial):Whenever you create a derived context that returns a cancel function,
 you must call that function, usually via defer. If you don't, the child context and its internal
 timers remain attached to the parent in memory until the parent dies, causing a memory and
 goroutine leak.
*/

func fetchWithTimeout() {
	// INFO: Creates a context that automatically cancels in 2 secs
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // INFO: always defer cancel to clean up resources

	// Pass the context to the worker
	result := make(chan string)
	go worker(ctx, result)
	select {
	case res := <-result:
		fmt.Println("Success: ", res)
	case <-ctx.Done():
		// ctx.Done() returns a channel that is closed when the context is canceled.
		// Reading from a closed channel unblocks immediately.
		fmt.Println("Failed: ", ctx.Err())
	}
}

func worker(ctx context.Context, res chan<- string) {
	select {
	case <-time.After(5 * time.Second):
		res <- "Data payload"
	case <-ctx.Done():
		fmt.Println("Worker abandoning task...")
		return
	}
}

var once sync.Once

func initConfig() {
	fmt.Println("Initializing...")
	// DEADLOCK! initConfig calls itself via the same sync.Once.
	// It is waiting for the mutex to unlock, but the mutex won't
	// unlock until initConfig finishes.
	once.Do(initConfig)
}

func main() {
	once.Do(initConfig)
	fetchWithTimeout()
}
