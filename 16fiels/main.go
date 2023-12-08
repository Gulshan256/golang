package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Working with files...")
	fmt.Println("======")

	contents := "This is a test file"

	files, error := os.Create("test.txt")
	CheckNillError(error)

	io.WriteString(files, contents)
	defer files.Close()

	readFile("test.txt")

}

func readFile(path string) (string, error) {
	file, error := os.ReadFile(path)
	CheckNillError(error)
	fmt.Println(string(file))

	return string(file), nil

}


func CheckNillError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
