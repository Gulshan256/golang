package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main() {
	fmt.Println("Hey welcome to the conversion program")
	fmt.Println("Rate your expressions from 1 to 10")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numRatings, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("Thanks for rating us ", numRatings)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRatings + 1)
		fmt.Printf("type: %T", numRatings)

	}

}


