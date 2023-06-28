package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func stressHandler(w http.ResponseWriter, r *http.Request) {
	// Get the number query parameter from the request
	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		http.Error(w, "Missing 'number' query parameter", http.StatusBadRequest)
		return
	}

	// Convert the number string to float64
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		http.Error(w, "Invalid 'number' query parameter", http.StatusBadRequest)
		return
	}

	// Get the number of CPUs available
	numCPUs := runtime.NumCPU()

	// Create a WaitGroup to synchronize Goroutines
	var wg sync.WaitGroup
	wg.Add(numCPUs)

	// Start the timer
	startTime := time.Now()

	// Simulate CPU-intensive workload using Goroutines
	for i := 0; i < numCPUs; i++ {
		go func() {
			// Perform some computationally expensive task
			calculateSquareRoot(number)
			wg.Done()
		}()
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Stop the timer and calculate the duration
	duration := time.Since(startTime)

	fmt.Fprintln(w, duration.String())
}

func calculateSquareRoot(number float64) {
	// Perform a computationally expensive task
	result := math.Sqrt(number)
	fmt.Printf("Square root of %.0f: %.2f\n", number, result)
}

func main() {
	// Print the current runtime architecture
	fmt.Println("Current architecture:", runtime.GOARCH)
	http.HandleFunc("/", handler)
	http.HandleFunc("/stress", stressHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
