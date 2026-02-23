package main

import "fmt"

// in case of an error Go executes all the values that were deferred

func main() {

	fmt.Println("counting...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	defer fmt.Println("worlddddddd") //get pushed to a stack LIFO
	defer fmt.Println("kiska")
	defer fmt.Println("hai ye")
	defer fmt.Println("tumko")
	fmt.Println("hullooo")
}
