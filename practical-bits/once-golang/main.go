package main

import (
	"sync"
	"time"
)

var shutdownOnce sync.Once

func Shutdown() {
	println("enter shutdown")
	shutdownOnce.Do(func() {
		// perform shutdown operations
		println(">>shutdown")
	})
}

func main() {
	// simulate multiple attempts to shutdown
	go Shutdown()
	go Shutdown()
	go Shutdown()

	// wait for goroutines to finish
	time.Sleep(1 * time.Second)
}
