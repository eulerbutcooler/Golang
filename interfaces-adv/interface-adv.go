package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// An interface that has a String method is called STRINGER
// Here the Person struct has a "String" method hence it overrides the internal representation of string that golang has
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur", 22}
	z := Person{"Ari", 23}
	fmt.Println(a, z)
}
