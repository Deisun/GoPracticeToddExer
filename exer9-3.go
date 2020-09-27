package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mux sync.Mutex

	counter := 0
	numGoroutines := 500
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			time.Sleep(time.Millisecond * 2)
			mux.Lock()
			counter++
			mux.Unlock()
			fmt.Printf("Goroutines: %v\n", runtime.NumGoroutine())
			wg.Done()
		}()

	}
	wg.Wait()

	fmt.Printf("Counter total: %v\n", counter)

}
