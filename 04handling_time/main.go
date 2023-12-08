package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello World")
	presentTime := time.Now()
	fmt.Println(presentTime)

	// change the format of the time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// create a custom time
	createdTime := time.Date(2023, time.November, 22, 15, 18, 118, 1, time.Now().Local().Location())
	fmt.Println(createdTime.Format("01-02-2006 15:04:05 Monday"))

}
