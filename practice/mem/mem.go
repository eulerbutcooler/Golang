package main

import (
	"bytes"
	"fmt"
	"sync"
)

/*
 * INFO: Escape Analysis (Stack vs. Heap)
Every Go application has two places to put variables: the Stack and the Heap.
The Stack: Insanely fast. Memory is automatically reclaimed the microsecond a function returns.
The Heap: Slower to allocate. Requires the Garbage Collector (GC) to track and clean it up, which
costs CPU cycles.
Escape Analysis is the compiler phase that determines where a variable lives.
Rule of Thumb: If a variable is only used inside the function that created it, it stays on the stack.
If a reference (pointer) to that variable is returned to the caller, the variable "escapes" to the heap
because it must outlive the function that created it.
*/

func main() {
	fmt.Println(onHeap())
	fmt.Println(onStack())
	logIt("PLUH")
}

func onStack() int {
	x := 10
	return x
}

func onHeap() *int {
	y := 20
	return &y
}

/*
 * INFO: sync.Pool
 sync.Pool is a thread-safe object recycler. Allocating memory on the heap is expensive.
 If you are constantly creating and destroying temporary objects
 like bytes.Buffer for incoming HTTP requests, you will choke the Garbage Collector.
*/
/*
  * INFO: sync.Pool lets you save those used objects and reuse them for the next request.
  Crucial Warning: sync.Pool is NOT a database connection pool or a cache. Why?
  Because the Garbage Collector is allowed to completely empty a sync.Pool at any time without warning.
  Items in a pool have no guaranteed lifespan. It is strictly for reusing memory allocations.
*/

var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer) // How to make a new one if the pool is empty
	},
}

func logIt(data string) {
	// 1. Get a buffer (might be fresh, might be recycled)
	buf := bufPool.Get().(*bytes.Buffer)

	// 2. Clear it out before use! (You don't know who used it last)
	buf.Reset()

	buf.WriteString(data)
	fmt.Println(buf.String())

	// 3. Put it back for someone else to use
	bufPool.Put(buf)
}
