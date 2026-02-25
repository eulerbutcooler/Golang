package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5} // len=5 cap=5
	b := a[1:3]               // len=2 cap=4
	b[0] = 99                 // modifies b but also modifies a
	fmt.Println(a)            // [1,99,3,4,5]

	b = append(b, 129) // b has cap=4 so it doesn't grow it overwrites
	fmt.Println(a)     // [1,99,3,129,5] <- 4 was overwritten
}
