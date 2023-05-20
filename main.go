package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/thecsw/haruhi"
	"github.com/thecsw/rei"
)

const (
	// defaultShortenerUrl is the default url for the shortener.
	defaultShortenerUrl = "https://photos.sandyuraz.com"
)

var (
	// shortenerUrl is the url for the shortener.
	shortenerUrl *string
	// createUrl is the url for the create endpoint.
	createUrl string
	// auth is the authentication token.
	auth *string
)

func main() {
	// Parse flags
	shortenerUrl = flag.String("shortener", defaultShortenerUrl, "monokuma shortener's url")
	urlsFile := flag.String("urls", "", "urls to shorten (newline-separated)")
	url := flag.String("url", "", "url to shorten")
	key := flag.String("key", "", "key to use when shortening a url")
	auth = flag.String("auth", "", "authentication token")
	flag.Parse()

	// createUrl is the url for the create endpoint.
	createUrl = *shortenerUrl + "/create"

	// Check if urls file was given
	if processUrls(*urlsFile) {
		return
	}

	// Check if url was given
	if len(*url) < 1 {
		log.Fatal("received an empty url")
	}

	// Get short url
	shortUrl, err := getShortUrl(*url, *key)
	if err != nil {
		log.Fatalf("shortening %s: %v\n", *url, err)
	}

	// Print short url
	fmt.Println(shortUrl)
}

// processUrls processes the urls in the given file.
func processUrls(urlsFile string) bool {
	// Check if urls file exists
	if !rei.FileMustExist(urlsFile) {
		return false
	}
	// Read urls file
	urlData, err := os.ReadFile(urlsFile)
	if err != nil {
		log.Fatalf("opening url file %s: %v", urlsFile, err)
	}
	// Split urls
	urls := strings.Split(string(urlData), "\n")
	// Process urls
	for _, url := range urls {
		// Skip empty urls
		if len(url) < 1 {
			continue
		}
		key := ""
		// See if url has a key
		if strings.Contains(url, ",") {
			// Split url and key
			urlAndKey := strings.Split(url, ",")
			url = urlAndKey[0]
			key = urlAndKey[1]
		}
		// Get short url
		shortUrl, err := getShortUrl(url, key)
		if err != nil {
			// Print error to stderr
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Printf("%s,%s\n", url, shortUrl)
	}
	return true
}

// getShortUrl returns the shortened url for the given url
// and key. If key is empty, it is not sent.
func getShortUrl(url, key string) (string, error) {
	// Build request
	request := haruhi.
		URL(createUrl).
		BodyString(url).
		Header("Authorization", "Bearer "+*auth)

	// Add key if present
	if len(key) > 0 {
		request = request.Param("key", key)
	}

	// Send request
	result, err := request.Post()

	// Check for errors
	if err != nil {
		return "", fmt.Errorf("shortening %s: %v", url, err)
	}
	return result, nil
}
