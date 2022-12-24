package main

import (
	"fmt"
	"flag"
	"strings"
	"os"
	"net/http"
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

	// Replace "example.com" with the URL you want to check
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the "Accept-Encoding" header to "gzip" to request gzip-compressed content
	req.Header.Set("Accept-Encoding", "gzip")

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the "Content-Encoding" header is set to "gzip"
	if resp.Header.Get("Content-Encoding") == "gzip" {
		fmt.Println("The URL:", url, "supports gzip compression")
	} else {
		fmt.Println("The URL:", url, "does not support gzip compression")
	}
}
