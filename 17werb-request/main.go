package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://lco.dev"

func main() {

	fmt.Println("Working with  web requests...")
	fmt.Println("======")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n", resp)
	defer resp.Body.Close()
	fmt.Println("Status code:", resp.StatusCode)

	// read response body
	bs, _ := io.ReadAll(resp.Body)
	fmt.Println("Response body:", string(bs))

}
