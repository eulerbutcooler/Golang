package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var c, java, python bool //all are of same type so we can define their type at the end -> "bool"
// c:=9 //not possible outside of a function
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%v>=%v\n", v, lim)
	}
	return lim
}
func main() {

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println(("tomorrow"))
	case today + 2:
		fmt.Println("In two days")
	}
	fmt.Println(pow(3, 2, 8), pow(3, 2, 20))

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Nice yar rich log")
	case "linux":
		fmt.Println("Tagda")
	default:
		fmt.Println("Abey jaa na")
	}

	i := 10
	// var k int //initializes k to 0 //can be used outside the function
	// var k = 0 //another method of declaring a variable
	// k:=0    // initializes k to 0
	fmt.Println("Hello, vurl", i, c, java, python)
	fmt.Println("Yo Mr. ", rand.Intn(12))
	fmt.Println("Yo Mr. ", math.Pi)
	fmt.Println(add(42, 42))
	a, b := swap("hello", "corn")
	fmt.Println(a, b)
	fmt.Println(split(28))
}

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	// returns x,y without explicitly mentioning them
	x = sum * 4 / 9
	y = sum - x
	return
}

// ALTERNATIVE***
// func split(sum int) (int, int) {
// 	x := sum * 4 / 9
// 	y := sum - x
// 	return x,y
// }
