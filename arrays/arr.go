package main

import "fmt"

func main() {
	var a [2]string //[2]string = number of values inside the array and their type
	a[0] = "huh"
	a[1] = "yeahhh"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} //sizes are part of type of array
	fmt.Println(primes)
}

// *** Note: Arrays can't be resized
