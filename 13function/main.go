package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Function in golang\n")
	fmt.Println("======")
	greeter("Rajeev")
	fmt.Print(proAdder(1, 2, 3, 4, 5))

}

func greeter(name string) {
	fmt.Println("Hey 😎", name)
}

func proAdder(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
