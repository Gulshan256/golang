package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// var name string
	// fmt.Println("Enter your name: ")
	// fmt.Scan(&name)
	// fmt.Println("Hello ", name)
	// fmt.Println("Enter your age: ")
	// var age int
	// fmt.Scan(&age)
	// fmt.Println("hey my name is ", name, " years old") 	

	//  
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ",)

	// comma ok idiom // error handling
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hello ", input)
		fmt.Printf("type of input is: %T", input)
	}





}