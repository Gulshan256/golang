package main

import (
	// "crypto/rand"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course -file
type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice float64 `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Website  string `json:"website"`
}

// Fake db
var courses []Course

// middleware/helpers -file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""

}

func main() {

	fmt.Println("Working with Build API...")

}

// Controller -file
// serve home routes

func severHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Home Page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	prams := mux.Vars(r)

	json.NewEncoder(w).Encode(courses)

	// loop through courses and return with matching id and course
	for _, c := range courses {
		if c.CourseID == prams["id"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with that id " + prams["id"])
	return

}



func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// what about data like that? {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data found")
		return
	}

	// generate unique id
	// append to courses

	course.CourseID = strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}


