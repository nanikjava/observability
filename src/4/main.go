package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numGoroutines = 15
	dataSize      = 1024 * 1024 // 1 MB
)

var (
	globalData []byte
	mutex      sync.Mutex
	wg         sync.WaitGroup
)

func appendData() {
	defer wg.Done()
	data := make([]byte, dataSize)

	mutex.Lock()
	globalData = append(globalData, data...)
	mutex.Unlock()
}

func main() {
	fmt.Println("Starting goroutines")

	// Loop to 1,000,000 before spinning up goroutines
	for i := 0; i < 1000000; i++ {
		// Spin up 1000 goroutines
		for j := 0; j < numGoroutines; j++ {
			wg.Add(1)
			go appendData()
		}

		// Wait for all goroutines to complete
		wg.Wait()

		// Sleep for 2 seconds
		time.Sleep(2 * time.Second)
	}

	fmt.Println("All goroutines finished")
	fmt.Printf("Total data size: %d MB\n", len(globalData)/(1024*1024))
}
