package main

import "fmt"

// Closure: When an outer function returns an inner function then that inner func
// will have access to the outer func variables even after the outer func returns
// and because of that the inner func can use that as a state across different func calls

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := range 10 {
		fmt.Println(pos(i), neg(-2*i))
	}
}
