package main

import (
	"fmt"
	"math"
)

// Since functions are first class citizens in golang we can assign a func to a variable
// These are called anonymous functions

// Here a function signature is used as a parameter of another function
// func(float64,float64)float64 matches the function signature of hypot function
// hence hypot can be passed to compute as a parameter
// similarly the math.Pow has the same func sign - func(float64, float64)float64
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(4, 5))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
