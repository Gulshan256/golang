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

	fmt.Println("======")
	fmt.Println(gulshan.GetAge())
	fmt.Println(gulshan.NewAge())

}

type Person struct {
	Name string
	Age  int
}

func (p Person) GetAge() string {
	return fmt.Sprintf("Age is %v", p.Age)
}

func (p Person) NewAge() int {
	return p.Age + 1
}
