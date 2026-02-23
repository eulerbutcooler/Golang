package main

import "fmt"

func IndexInt(s []int, x int) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func IndexStr(s []string, x string) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// Generic function that does the job of both IndexInt and IndexStr
// Here T is a parameter and comparable means that it must be comparable hence == operations can be done on it
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 30, -15, 11}
	fmt.Println(IndexInt(si, -15))
	ss := []string{"hello", "mello", "dirty", "fellow"}
	fmt.Println(IndexStr(ss, "yo"))
	// --------------------------------------------------
	fmt.Println(Index(si, 90))
	fmt.Println(Index(ss, "dirty"))
}
