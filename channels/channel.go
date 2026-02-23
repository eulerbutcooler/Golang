package main

import "fmt"

// Channels are container datatypes that allow two go routines to communicate with each other
// they are like tunnels by which go routines can talk to each other

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to c
}

// -----------------------------------
// Example showing infinite loop and switch case
func fibselect(c, quit chan int) {
	x, y := 0, 1
	// for loop without conditions is an infinite loop
	for {
		select {
		case c <- x: // if we can insert x in channel c then it'll be true and executed
			x, y = y, x+y
		case <-quit: // if there's any value coming from quit channel this will execute
			fmt.Println("quit")
			return
		}
	}
}

// -----------------------------------

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for _ = range n {
		c <- x
		x, y = y, x+y
	}
	close(c) // to inform the main routine that we are done with this goroutine
}

func main() {
	cfbs := make(chan int)
	quit := make(chan int)
	// anonymouse function
	go func() {
		// this loops over channel cfbs and receives from it in each iteration
		for _ = range 10 {
			fmt.Println(<-cfbs)
		}
		// after we are done from 0 to 9 we send a value to "quit" channel
		quit <- 0
	}()
	// this is called from the main go routine
	fibselect(cfbs, quit)

	// ----------------------------------
	cfb := make(chan int, 10)
	go fibonacci(cap(cfb), cfb)
	// when we range over a channel it serially takes out the values from itself as
	// other go routine is appending the values
	for i := range cfb {
		fmt.Println(i)
	}

	// ----------------------------------
	s := []int{10, 20, 30, 40, -1, 2, 9, 8}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //receive from c
	fmt.Println(x, y, x+y)
	// --------------------------------
	ch := make(chan int, 2) //buffered channel
	ch <- 1
	ch <- 2
	// ch <- 3 //causes a deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}

// Two types of channels - buffered and unbuffered
// Buffered channels are like queues
// Go routines can push and pull from the buffered channels until their capacity is full
// Producer consumer thingie
// ---------------------------------
// Unbuffered channels get blocked unless both the consumer and producer are in sync
// These are usually used for synchronizing/blocking
