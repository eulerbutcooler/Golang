package main

import "fmt"

// Maps = Hashmaps.
// Maps should be used when we have to find something quickly
// The keys should be comparable, rest, we can store anything
// Lookup, Adding, Deleting time of maps = O(1)
// Struct are for efficiently storing values

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex //[string]=type of key, Vertex=type of value

func main() {
	// make syntax for initializing maps
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		46.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	hsh := map[string]Vertex{
		"Google": Vertex{
			37.42202, -122.0848,
		},
	}
	fmt.Println(hsh)

	// We can only define maps with either literal syntax or make syntax.
	// var syntax created maps can't have keys and values added to them as they
	// are initialized as nil
	// so either use - m:=make(map[string]int)
	// or - m:=map[string]int{}

	ml := make(map[string]int)
	ml["Answer"] = 43
	fmt.Println("The value is: ", ml["Answer"])
	ml["Answer"] = 48
	fmt.Println("The value is: ", ml["Answer"])
	ml["Wrong"] = 455
	fmt.Println("The value is: ", ml["Wrong"])

	for key, value := range ml {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	delete(ml, "Answer")
	fmt.Println("The value: ", ml["Answer"])
	// To differentiate between absent value and zero value
	// here ok = bool
	v, ok := ml["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)
}

// When to use a map or a struct?
// Can't loop over a struct as it is a static type
// Maps' values dont have a fixed memory location and hence we cant use pointers
// Structs' values have a fixed memory location and hence we can use pointers
// When we loop over maps the keys are in random order and we explicitly have to sort
// them out if we want to use keys in a specific order
// Mental model - Structs for storing things
// 				  Maps for looking up things
