package main

import (
	"fmt"
)

func main() {
	fmt.Println("Struct in golang")

	fmt.Println("======")
	gulshan := Person{"Gulshan", 25}
	fmt.Println(gulshan)
	fmt.Printf("Details is : %+v\n", gulshan)
	fmt.Printf("Name is : %v\n", gulshan.Name)

}

type Person struct {
	Name string
	Age  int
}
