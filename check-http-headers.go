package main

import (
	"fmt"
	"flag"
	"strings"
	"net/http"
	"os"
)

func main() {

	// Define a string flag "-u"/"--url" which needs to be provided to program
	var url string
	flag.StringVar(&url, "u", "", "URL flag")
	flag.StringVar(&url, "url", "", "URL flag")

	// Parse the command-line flags
	flag.Parse()

	// Check if the variable does not start with "http://" or "https://"
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		fmt.Println("Variable does not contain a URL with the \"http\" or \"https\" scheme")
		os.Exit(1)
	}
	
	// Check if a URL was provided as an argument
	/* disabled by anish
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run headers.go URL")
		os.Exit(1)
	}
	*/

	// Make an HTTP GET request to the provided URL
	//response, err := http.Get(os.Args[1])
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	// Print the headers
	fmt.Println("HTTP Headers:")
	for key, values := range response.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}

