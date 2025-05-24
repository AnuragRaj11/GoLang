package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("url package main function")

	rawURL := "https://example.com/path?query=123"
	fmt.Println("URL:", rawURL)
	fmt.Println("URL Length:", len(rawURL))

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Parsed URL:", parsedURL)
	fmt.Println("URL Scheme:", parsedURL.Scheme)
	fmt.Println("URL Host:", parsedURL.Host)
	fmt.Println("URL Path:", parsedURL.Path)
	fmt.Println("URL Query:", parsedURL.Query())

	fmt.Println("URL String:", parsedURL.String())

	jsonData, err := json.Marshal(parsedURL)
	if err != nil {
		fmt.Println("Error marshaling URL to JSON:", err)
		return
	}
	fmt.Println("JSON representation of URL:", string(jsonData))
}
  