package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://example.com:8080/path?query=param#fragment"
	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		fmt.Println("Error parsing URL: ", err)
	}

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Port: ", parsedURL.Port())
	fmt.Println("Path: ", parsedURL.Path)

	rawURL1 := "https://example.com/finder?name=John&age=30"
	parsedURL1, err := url.Parse(rawURL1)

	if err != nil {
		fmt.Println("Error parsing url 1: ", err)
		return
	}

	queryParams := parsedURL1.Query()
	fmt.Println("Query Params: ", queryParams)

	fmt.Println("Name: ", queryParams.Get("name"))
	fmt.Println("Age: ", queryParams.Get("age"))

	//building a new url ourself

	baseURL := &url.URL{
		Scheme: "https",
		Host: "razaprinters.com",
		Path: "/invites",
	}

	query := baseURL.Query()
	query.Set("name", "John")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Generated URL: " , baseURL.String())

	//another way to build url

	values := url.Values{}

	values.Add("name", "John Doe")
	values.Add("age", "30")
	values.Add("city", "Gothenburg")
	values.Add("country", "Sweden")

	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	newBaseUrl := "https://facebook.com/search"
	fullUrl := newBaseUrl + "?" + encodedQuery

	fmt.Println(fullUrl)
	
}