package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/gorilla/mux"
)

// Struct to store crawling result
type CrawlingResult struct {
	URL     string
	Content string
}

func crawlURL(url string, userAgent string, resultChan chan<- CrawlingResult, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a new context for each crawling URL
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set User-Agent header in the request
	err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		// Set the User-Agent header
		headers := map[string]interface{}{
			"User-Agent": userAgent,
		}
		err := network.SetExtraHTTPHeaders(headers).Do(ctx)
		if err != nil {
			return err
		}
		return nil
	}))

	if err != nil {
		log.Printf("Error setting User-Agent for URL %s: %v\n", url, err)
		return
	}

	// Navigate to the URL
	err = chromedp.Run(ctx, chromedp.Navigate(url))
	if err != nil {
		log.Printf("Error navigating to URL %s: %v\n", url, err)
		return
	}

	// Wait for the page to load completely
	err = chromedp.Run(ctx, chromedp.WaitReady("body"))
	if err != nil {
		log.Printf("Error waiting for page to load for URL %s: %v\n", url, err)
		return
	}

	// Get the full HTML content of the page
	var htmlContent string
	err = chromedp.Run(ctx, chromedp.OuterHTML("html", &htmlContent))
	if err != nil {
		log.Printf("Error getting HTML content for URL %s: %v\n", url, err)
		return
	}

	// Send the result to the channel
	resultChan <- CrawlingResult{URL: url, Content: htmlContent}
}

func crawlHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the target URL from the query parameter
	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		http.Error(w, "URL target must be provided, example: http://localhost:8080/crawl?url=https://cmlabs.co", http.StatusBadRequest)
		return
	}

	// Get the User-Agent from the query parameter (use default if empty)
	userAgent := r.URL.Query().Get("user_agent")
	if userAgent == "" {
		userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
	}

	// Split targetURL by comma to get multiple URLs if any
	urls := strings.Split(targetURL, ",")

	// Create a channel to receive crawling results
	resultChan := make(chan CrawlingResult, len(urls))

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start concurrent crawling using goroutine
	for _, url := range urls {
		wg.Add(1)
		go crawlURL(url, userAgent, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	// Handle results
	for result := range resultChan {
		filePath := fmt.Sprintf("output_%s.html", strings.ReplaceAll(result.URL, "https://", ""))
		file, err := os.Create(filePath)
		if err != nil {
			log.Printf("Error creating file for URL %s: %v\n", result.URL, err)
			continue
		}

		_, err = file.WriteString(result.Content)
		if err != nil {
			log.Printf("Error writing content for URL %s: %v\n", result.URL, err)
		}

		file.Close()
	}

	fmt.Fprintln(w, "Crawling finished. Results are saved in HTML files.")
}

func main() {
	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Define the handler for the /crawl endpoint
	router.HandleFunc("/crawl", crawlHandler)

	// Start the server on port 8080
	fmt.Println("API for website crawling is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
