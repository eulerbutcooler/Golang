package main

import "fmt"

/*
 *  INFO: Iska scene undefined hai matlab kuch bhi ho sakta hai message display bhi ho sakta hai
 * ya fir infinite loop mein bhi atak sakte hai kyuki "happens-before" wala concept is not defined yahan.
 * usko define karne ke liye we'll use a channel
 */

var msg string
var done bool

func setup() {
	msg = "yo pierre"
	done = true
}

var ch = make(chan struct{})

/*
- NOTE: Ismein we are using a channel and this way we ensure happens-before
- guarantee aur hello will always be printed
- execution order is guaranteed to be:
- 1. `msg = "ji mein oppressed majority"` — writes the message
- 2. `ch <- struct{}{}` — sends the signal (this **happens-before**...)
- 3. `<-ch` — ...the receive completes in `main`
- 4. `fmt.Println(msg)` — reads the message (guaranteed to see the write from step 1)
*/
func setupwithappensbefore() {
	msg = "ji mein oppressed majority"
	ch <- struct{}{}
}

func main() {
	go setup()
	for !done {
	}
	fmt.Println(msg)

	go setupwithappensbefore()
	<-ch
	fmt.Println(msg)
}
