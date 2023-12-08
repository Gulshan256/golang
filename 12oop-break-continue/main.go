package main

import (
	"fmt"
)

func main() {

	fmt.Println("Loop, break and continue in golang")
	fmt.Println("======")

	// for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// for i, day := range days {
	// 	fmt.Printf("Index: %d, Day: %s\n", i, day)
	// }

	for _, day := range days {
		fmt.Printf("Day: %s\n", day)
	}

	rougeValue := 1

	for rougeValue < 10 {
		fmt.Println(rougeValue)
		rougeValue++

		if rougeValue == 5 {
			rougeValue++ //
			continue
		}

		if rougeValue == 8 {
			break
		}

	}

}
