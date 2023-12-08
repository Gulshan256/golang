package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	// create a pointer to a string
	var firstName *string = new(string)
	fmt.Println("firstName: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	*firstName = input


	fmt.Println(firstName)  // print the memory address
	fmt.Println(*firstName) // print the value at the memory address

}
