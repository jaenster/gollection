# Gollection!

Just a simple lib for collections in go.

Looks a bit like a set in javascript

```go

package main

import (
	"github.com/jaenster/gollection"
)

type MyStruct struct {
	
}

func main() {
	one := &MyStruct{}
	two := &MyStruct{}

	gollection := gollection.New[MyStruct]()

	gollection.Add(one)
	fmt.Printf("size %d", gollection.Size()) // 1

	gollection.Add(one)                      // Doesnt get added again
	fmt.Printf("size %d", gollection.Size()) // 1

	gollection.Add(two)
	fmt.Printf("size %d", gollection.Size()) // 2

	count := 0
	gollection.ForEach(func(ms *MyStruct) {
		count++
	})
	fmt.Printf("size %d", count) // 2

	gollection.Remove(two)
	fmt.Printf("size %d", count) // 1
}

```