package main

import (
	"fmt"
	"math"
)

// Interface = contains method signature
// Interface gives us a signature that a type should have to implement that interface
// As long as there is a type and it has a method whose signature and whose receiver
// matches that of the interface then we can say that this type implements this interface (Abser)
type Abser interface {
	Abs() float64
}

type Myfloat float64
type Vertex struct {
	X, Y float64
}

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// -------------

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we dont need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var a Abser
	f := Myfloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f  // a Myfloat implements Abser
	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
	// --------------------
	var i I = T{"hullo"}
	i.M()
	// --------------------
	var it interface{} = "hello" //empty interface. can also use "any" instead of interface here
	s := it.(string)
	fmt.Println(s)

	s, ok := it.(string)
	fmt.Println(s, ok)

	fl, ok := it.(float64)
	fmt.Println(fl, ok)

	// fl1 := it.(float64) //panic. we get a runtime error
	// fmt.Println(fl1)

	// ------------------
	do(21)
	do("huh")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%s is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I dont know about type %T\n", v)
	}
}
