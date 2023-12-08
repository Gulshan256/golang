package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Working with Modules...")
	fmt.Println("======")
	greeting()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))



}

func greeting() string {
	return "Hello, there mod users!"

}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Welcome to the home page!"))

}
