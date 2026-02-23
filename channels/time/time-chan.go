package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v map[string]int
}

func (c *Counter) Inc(key string) {
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	return c.v[key]
}

func main() {
	// c := Counter{v: make(map[string]int)}
	// this spawns a 1000 go routines and that causes concurrent map writes - error
	// we'll have to use mutex
	// for _ = range 1000 {
	// 	go c.Inc("somekey")
	// }
	// time.Sleep(time.Second)
	// fmt.Println(c.Value("somekey"))
	// ---------------------------------------
	sc := SafeCounter{v: make(map[string]int)}
	for _ = range 1000 {
		go sc.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(sc.Value("somekey"))
}

// Safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// lock so only one goroutine at a time can access the map c.v
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v
	// defer unlocks the mutex when we return from this function
	defer c.mu.Unlock()
	return c.v[key]
}
