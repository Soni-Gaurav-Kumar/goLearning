package main

import (
	"fmt"
	"net/url"
)

func Url() {
	// Parse a URL string in to a URL Object
	myUrl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"

	parsedURL, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println("error while Parsing URL:", err)
		return
	}

	// Accessing URL Components
	fmt.Println("URL Scheme :", parsedURL.Scheme)
	fmt.Println("URL Host :", parsedURL.Host)
	fmt.Println("URL path :", parsedURL.Path)
	fmt.Println("URL RawQuery :", parsedURL.RawQuery)

	// Modifying URL Components
	parsedURL.Scheme = "http"
	parsedURL.Host = "Soni.com"
	parsedURL.Path = "/gaurav/kumar/soni"
	parsedURL.RawQuery = "id=609&&team=uplifters"

	// Constructing a URL string from URL object
	newUrl := parsedURL.String()
	fmt.Println("Modified URL : ", newUrl)

}
