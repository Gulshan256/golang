package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	// "strings"
)

func main() {
	fmt.Println("Working with Request Verb...")

	// perfomeGetRequest("http://127.0.0.1:5000/api/data")

	// perfomePostRequest("http://127.0.0.1:5000/api/data")
	perfomePostFormRequest()

}

// func perfomeGetRequest(url string) {

// 	requestURL := url

// 	reapond, err := http.Get(requestURL)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer reapond.Body.Close()

// 	fmt.Println("Status Code:", reapond.StatusCode)
// 	fmt.Println("Content length:", reapond.ContentLength)
// 	fmt.Println("Content Type:", reapond.Header.Get("Content-Type"))
// 	content, _ := io.ReadAll(reapond.Body)
// 	fmt.Println("Content:", string(content))

// 	// // another way to get the content
// 	// var reapondContent strings.Builder
// 	// contents, _ := io.ReadAll(reapond.Body)
// 	// byteCount, _ := reapondContent.Write(contents)

// 	// fmt.Println("Content:", byteCount)

// }

func perfomePostRequest(url string) {

	requestURL := strings.TrimSpace(url)

	// Fake json data
	requestBody := strings.NewReader(
		`{
	
		"1": {"name": "John", "age": 30},
		"2": {"name": "Jane", "age": 25},
		"3": {"name": "Jack", "age": 20}  

		}`)

	reapond, err := http.Post(requestURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer reapond.Body.Close()

	fmt.Println("Status Code:", reapond.StatusCode)
	fmt.Println("Content length:", reapond.ContentLength)
	fmt.Println("Content Type:", reapond.Header.Get("Content-Type"))
	content, _ := io.ReadAll(reapond.Body)
	fmt.Println("Content:", string(content))

}

func perfomePostFormRequest() {

	const requestURL = "http://localhost:5000/api/form-data"

	// form data
	formData := url.Values{}

	formData.Add("name", "John")
	formData.Add("age", "30")

	response, err := http.PostForm(requestURL, formData)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	contant, _ := io.ReadAll(response.Body)

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	fmt.Println("Content:", string(contant))

}
