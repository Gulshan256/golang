// How to create JSON data in golang
package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
)

type JSON struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Id   int    `json:"id"`
}

func main() {
	fmt.Println("Working with JSON...")

	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {

	info := []JSON{
		{Name: "John", Age: 21, Id: 1},
		{Name: "Jane", Age: 22, Id: 2},
		{Name: "Joe", Age: 23, Id: 3},
		{Name: "Joe", Age: 23, Id: 3},
	}

	final, err := json.MarshalIndent(info, "", " \t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(final))

}

func DecodeJSON() {

	jsonData, err := json.Marshal(JSON{Name: "John", Age: 21, Id: 1})
	if err != nil {
		panic(err)
	}

	var final JSON

	if json.Valid(jsonData) {
		err = json.Unmarshal(jsonData, &final)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(final)

	// some case add data to key value pairs

	var final2 map[string]interface{}
	json.Unmarshal(jsonData, &final2)
	fmt.Printf("%+v\n", final2)

}
