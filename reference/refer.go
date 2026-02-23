package main

import "fmt"

// POINTERS

func main() {
	i, j := 42, 2071
	p := &i         //point to i
	fmt.Println(*p) //read i through the pointer
	*p = 87         //set i through the pointer
	fmt.Println(i)  //see new value of i

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

// *** Zero value of reference type is nil
