package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Methods = golang allows us to assign functions to types. any types.
// Can be struct can be integer can be string
// This is sorta like string methods in pythong s.slice(2:6), etc ig
// Methods have a receiver parameter (v Vertex) before the method name (Abs)
// The receiver is like "self" in python - it represents the instance the method is called on
func (v Vertex) Abs() float64 { //(v Vertex) is called the receiver
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Absfunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// receiver parameter can also be a pointer
// so whatever change we make here will reflect on the original instance
// Rule of thumb: Use pointer receivers for methods that modify the receiver or if the type is large (to avoid copying).
// If you've made one method pointer receiver make all other methods pointer receivers
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func Scalefunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	m := Myfloat(-math.Sqrt(123))
	// ----------
	v := Vertex{3, 4}
	v.Scale(2) //Dont have to explicitly do (&v).Scale(2). Golang automatically passed the pointer
	Scalefunc(&v, 10)

	p := Vertex{4, 3}
	p.Scale(3)
	Scalefunc(&p, 8)
	fmt.Println(m.Abs())
	fmt.Println(Absfunc(v))

	fmt.Println(v.Abs())
}

// This is analogous to `v.abs()` in Python if `Vertex` were a class with an `abs` method.
// The method operates on the data in `v` (X=3, Y=4) and returns the result.
