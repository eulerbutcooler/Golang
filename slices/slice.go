package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s[2])

	names := [4]string{
		"Desmond",
		"Molly",
		"Jones",
		"Happy",
	}
	a := names[0:2] //last index is not included in slices
	b := names[1:4]
	b[2] = "Ho chi minh"
	fmt.Println(a, b)
	fmt.Println(names)
	// ----------------------------------
	q := []int{2, 3, 4, 7, 11, 13} // we are creating a slice without a backing array here. go creates it for us in the background
	fmt.Println(q)
	r := []bool{true, false, false, true, false, true} //same here. hence we haven't specified the length
	fmt.Println(r)

	h := []struct { //slice of a struct
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, false},
		{11, true},
		{13, false},
	}
	fmt.Println(h)

	// ------------------------------

	li := []int{2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(li)

	li = li[:5]
	printSlice(li)

	li = li[0:2] //slice from a slice
	printSlice(li)

	// ------------------------------

	li = li[2:]
	printSlice(li)

	sl := make([]int, 5)
	printSlice(sl)

	// make = make(type,length,capacity) - when capacity is not provided cap=len
	// also length can never be greater than capacity
	sl2 := make([]int, 0, 5)
	printSlice(sl2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) //capacity = cap = first element that the slice points to till the end.
}
