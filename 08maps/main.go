package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")
	languages := make(map[string]string)
	languages["Java"] = "Static"
	languages["Javascript"] = "Dynamic"
	languages["Go"] = "Static"
	languages["Rust"] = "Static"
	languages["Python"] = "Dynamic"
	languages["C"] = "Static"
	languages["C++"] = "Static"

	fmt.Println("languages: ", languages)

	// delete
	delete(languages, "C++")
	fmt.Println("languages: ", languages)

	// loop over languages
	for key, val := range languages {
		fmt.Println("key: ", key, "val: ", val)
	}
}
