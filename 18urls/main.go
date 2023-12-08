package main

import (
	"fmt"
	"net/url"
)

const URL = "https://www.google.com/search?q=golang"

func main() {
	fmt.Println("Working with Handeling URL...")
	fmt.Println("======")

	// URL Parsing
	// fmt.Println("URL Parsing")
	// fmt.Println("======")
	// // Parse the URL and ensure there are no errors.
	// u, err := url.Parse(URL)
	// if err != nil {
	// 	panic(err)
	// }

	// // Print out the URL and the query string.
	// fmt.Println("URL:", u.String())
	// fmt.Println("Query:", u.Query())
	// fmt.Println("Scheme:", u.Scheme)
	// fmt.Println("Port:", u.Port())
	// fmt.Println("Path:", u.Path)
	// fmt.Println("Host:", u.Host)

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "www.google.com",
		Path:   "/search",
		RawQuery: url.Values{
			"q": {"golang"},
		}.Encode(),
	}

	fmt.Println(partsOfUrl.String())

}
