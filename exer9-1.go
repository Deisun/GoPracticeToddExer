package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("foo running...")
	wg.Done()
}
func bar() {
	fmt.Println("bar running...")
	wg.Done()
}
func main() {

	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
}
