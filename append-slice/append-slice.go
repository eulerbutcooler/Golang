package main

import "fmt"

// append is a variadic function in golang i.e. it can take up any number of args
// append will double the capacity everytime there is no space upto 256
func main() {
	var s []int
	s = append(s, 0)
	s = append(s, 1)
	s = append(s, 2, 3, 4, 5, 6, 7, 8, 9)
	printSlice("s", s)

	var nums []int
	for i := range 30 {
		nums = append(nums, i)
		fmt.Printf("Length: %v, Capacity: %v\n", len(nums), cap(nums))
	}

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}

func printSlice(x string, s []int) {
	fmt.Printf("%s - len=%v, cap=%v\n %v\n", x, len(s), cap(s), s)
}
