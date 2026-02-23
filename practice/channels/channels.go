package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		 * INFO: hum yahan buffered channel use karenge
			- syntax = ch:=make(chan [type] [size])
			- Sender only blocks if the capacity of the buffer is full
			- Receiver only blocks if the buffer is empty
	*/
	ch := make(chan string, 2)
	// INFO: Send hua in the buffer
	ch <- "yo"
	ch <- "bhaiji"
	fmt.Println("Pushed two words in chan")

	// INFO: Receive hua from buffer

	fmt.Println("Reading: ", <-ch)
	fmt.Println("Reading: ", <-ch)

	// INFO: sender blocks until receiver is present
	unbufch := make(chan string)
	go func() {
		fmt.Println("Sender: Koi toh lelo isey")
		// HACK: blocks here until main() reaches <-unbufch
		unbufch <- "Lao dedo"
		fmt.Println("Sender: Bhala hua mori gagri phuti")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Receiver: Bhejo yar")
	// HACK: yahan pe we received it
	msg := <-unbufch
	fmt.Println("Receiver: Milgayoo: ", msg)
}
