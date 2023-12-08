package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	// var name = []string{"John", "Paul", "George", "Ringo"}

	// fmt.Println(name)

	// name = append(name, "Pete", "John")

	// fmt.Println(name)

	// // slice form 1 slice to 3
	// fmt.Println(name[1])

	// // get last element
	// fmt.Println(name[len(name)-1])

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867

	highScore = append(highScore, 555, 666, 777)

	fmt.Println("len of highScore: ", len(highScore))

	// method in slice

	// // 1. copy
	// copySlice := make([]int, len(highScore))
	// copy(copySlice, highScore)

	// fmt.Println("copySlice: ", copySlice)

	//  sort
	sort.Ints(highScore)
	fmt.Println("highScore: ", highScore)

	language := []string{"C", "C++", "Java", "Python", "Go", "Rust", "Javascript"}
	fmt.Println("language: ", language)
	// how to remove element based on index
	index := 2
	language = append(language[:index], language[index+1:]...)
	fmt.Println("language: ", language)

}
