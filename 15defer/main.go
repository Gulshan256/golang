package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer in golang")

	defer fmt.Println("1")
	fmt.Println("======")

}
