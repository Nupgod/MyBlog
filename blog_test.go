package ex01

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func TestRateLimit(t *testing.T) {
	var wg sync.WaitGroup
	var tooManyRequestsCount int
	var okCount int

	// Function to make requests
	makeRequest := func(url string) {
		defer wg.Done()
		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("Error making request: %v", err)
			return
		}
		defer resp.Body.Close()

		// Increment the appropriate counter based on the response status code
		if resp.StatusCode == http.StatusTooManyRequests {
			tooManyRequestsCount++
		} else if resp.StatusCode == http.StatusOK {
			okCount++
		}
	}

	// Make 110 requests concurrently
	for i := 0; i < 150; i++ {
		wg.Add(1)
		go makeRequest("http://localhost:8888/blog/")
	}

	// Wait for all requests to finish
	wg.Wait()

	fmt.Printf("Ok Status total: %d\nManyReq status total: %d\n", okCount, tooManyRequestsCount)
}
