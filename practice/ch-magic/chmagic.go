package main

import "fmt"

func main() {
	selectfairness()
	closingchannels()
	nilchannels()
}

func selectfairness() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	/* INFO: Both channels are ready to be read.
	- When run multiple times, the output will randomly flip between 1 and 2.
	- This shows the fairness of the select statement. When both cases are valid at the same time,
	  it randomly selects one of them to execute ensuring that no single case is favored over the other.
	*/
	select {
	case val := <-ch1:
		fmt.Println("Read from ch1:", val)
	case val := <-ch2:
		fmt.Println("Read from ch2:", val)
	}
}

func closingchannels() {
	/*
		 * INFO: Only the sender should close the channel if the receiver closes it wont know if the sender
			* is done with sending or will be sending more.
	*/
	ch := make(chan int, 2)
	ch <- 243
	close(ch)
	/*
		 * INFO: Yahan ok se pata chaljayega ki the channel is closed or not
			* we cant trust the val always kyuki what if the channel had a 0 value?
			* we cant differentiate channel is closed or not based on the value
			* ok is a bool and it is not being redeclared it is only be reused
	*/
	val1, ok := <-ch
	fmt.Println(val1, ok)
	// 243, true
	val2, ok := <-ch
	// 0, false
	fmt.Println(val2, ok)
}

func nilchannels() {
	/*
		* WARN: If you declare a channel but don't initialize it with make(), it is nil.
		- Sending to a nil channel blocks forever.
		- Receiving from a nil channel blocks forever.
		- Closing a nil channel panics.
	*/
	/*
		 * INFO: When reading from two channels in a select loop, one channel closes then the select will
			* start infinitely looping on the closed channel (because receiving from a closed channel never
			* blocks), to stop this we make the channel nil
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)
	/* INFO: If both ch1 and ch2 become nil, this select blocks forever (deadlock).
	   Usually, you'd add a breakout condition here.
	*/
	for {
		// Breakout condition
		if ch1 == nil && ch2 == nil {
			fmt.Println("Both channels are closed. Exiting")
			return
		}
		select {
		case val, ok := <-ch1:
			if !ok {
				ch1 = nil
				// agar isko nil nahi kiya toh infinite loop mein chale jayenge
				continue
			}
			fmt.Println("Got from 1: ", val)
		case val, ok := <-ch2:
			if !ok {
				ch2 = nil
				// agar isko nil nahi kiya toh infinite loop mein chale jayenge
				continue
			}
			fmt.Println("Got from 2: ", val)
		}
	}
}
