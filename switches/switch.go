package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Goodmorning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

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

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Nice yar rich log")
	case "linux":
		fmt.Println("Tagda")
	default:
		fmt.Println("Abey jaa na")
	}
}
