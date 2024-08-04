package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func fetchJoke() (Joke, error) {
	resp, err := http.Get("https://official-joke-api.appspot.com/jokes/random")
	if err != nil {
		return Joke{}, err
	}
	defer resp.Body.Close()

	var joke Joke
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return Joke{}, err
	}

	return joke, nil
}

func fetchAndPrintJoke(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	joke, err := fetchJoke()
	if err != nil {
		fmt.Println("Error fetching joke:", err)
		return
	}

	jokePretty, err := json.MarshalIndent(joke, "", "  ")
	if err != nil {
		fmt.Println("Error formatting joke:", err)
		return
	}

	fmt.Println(string(jokePretty))
}

func main() {
	for {
		// Create a WaitGroup to wait for all goroutines to finish
		var wg sync.WaitGroup
		numGoroutines := 10

		// Spin up 10 goroutines
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1) // Increment the WaitGroup counter
			go fetchAndPrintJoke(&wg)
		}

		// Wait for all goroutines to finish
		wg.Wait()

		// Sleep for 10 seconds before starting the process again
		fmt.Println("All jokes fetched. Sleeping for 10 seconds...")
		time.Sleep(10 * time.Second)
	}
}
