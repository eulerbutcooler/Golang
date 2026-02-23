package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for _ = range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// this spawns a new goroutine that runs concurrently
	// this is managed by the go scheduler and not the OS
	// go routines are very light-weight. they are lighter than threads.
	// go routines run concurrently with the main go routine
	// the main go routine waits for the other go routines to finish
	go say("world")
	say("helloooo")
}
