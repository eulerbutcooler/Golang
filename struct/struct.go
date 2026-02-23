package main

import "fmt"

type Vertex struct {
	X int //we have to capitalize the field names inside of structs to expose them outside the package
	Y int
}

var (
	v1 = Vertex{1, 2}  //has type Vertex
	v2 = Vertex{X: 5}  //Y:0 is implicit
	v3 = Vertex{}      //X:0 and Y:0
	p  = &Vertex{1, 2} //has type *Vertex
)

func main() {

	fmt.Println(v1, v2, v3, p)

	fmt.Println(Vertex{1, 2})

	v := Vertex{22, 23}
	v.X = 55 //accessing items
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Println(v)
}
