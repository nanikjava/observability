package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type DogImageResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Define Prometheus metrics
var (
	imagesFetched = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "dog_images_fetched_total",
			Help: "Total number of dog images fetched.",
		},
	)
	fetchDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "dog_image_fetch_duration_seconds",
			Help:    "Histogram of the duration of dog image fetches.",
			Buckets: prometheus.DefBuckets,
		},
	)
	httpStatusCodes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "dog_http_status_codes_total",
			Help: "Total number of HTTP status codes received.",
		},
		[]string{"status_code"},
	)
	fetchErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "dog_images_fetched_errors_total",
			Help: "Total number of errors when trying to fetch images.",
		},
	)
)

func init() {
	// Register metrics with Prometheus
	prometheus.MustRegister(imagesFetched)
	prometheus.MustRegister(fetchDuration)
	prometheus.MustRegister(httpStatusCodes)
	prometheus.MustRegister(fetchErrors)
}

func fetchImageURL() (string, error) {
	start := time.Now()
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Record the HTTP status code
	httpStatusCodes.WithLabelValues(fmt.Sprintf("%d", resp.StatusCode)).Inc()

	var dogImageResp DogImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&dogImageResp); err != nil {
		return "", err
	}

	if dogImageResp.Status != "success" {
		return "", fmt.Errorf("unsuccessful status: %s", dogImageResp.Status)
	}

	// Record the duration
	fetchDuration.Observe(time.Since(start).Seconds())
	imagesFetched.Inc() // Increment the images fetched counter

	return dogImageResp.Message, nil
}
func downloadImage(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading image:", err)
		fetchErrors.Inc()
		return
	}
	defer resp.Body.Close()

	// Copy the image data into memory
	var imageBuffer bytes.Buffer
	_, err = io.Copy(&imageBuffer, resp.Body)
	if err != nil {
		fmt.Println("Error copying image data:", err)
		return
	}

	// Process the image data in memory
	fmt.Printf("Downloaded image size: %d bytes\n", imageBuffer.Len())
}

func fetchAndDownloadImage(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	imageURL, err := fetchImageURL()
	if err != nil {
		fetchErrors.Inc()
		fmt.Println("Error fetching image URL:", err)
		return
	}

	downloadImage(imageURL, wg)
}

func main() {
	// Start Prometheus HTTP handler
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Println("Starting Prometheus metrics server on :2112")
		if err := http.ListenAndServe(":2112", nil); err != nil {
			fmt.Println("Error starting Prometheus metrics server:", err)
		}
	}()

	for {
		// Create a WaitGroup to wait for all goroutines to finish
		var wg sync.WaitGroup
		numGoroutines := 2

		// Spin up 10 goroutines
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1) // Increment the WaitGroup counter
			go fetchAndDownloadImage(&wg)
		}

		// Wait for all goroutines to finish
		wg.Wait()

		// Sleep for 10 seconds before starting the process again
		fmt.Println("All images fetched and downloaded. Sleeping for 10 seconds...")
		time.Sleep(10 * time.Second)
	}
}
