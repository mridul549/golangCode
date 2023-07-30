package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// create a modal for courses - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// db
var courses []Course

// middlewares or helpers - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	// seeding
	courses = append(courses, 
		Course{CourseId: "1", CourseName: "react", CoursePrice: 499, Author: &Author{Fullname: "Antonio Bowman", Website: "http://deg.bd/dagrejig"}},
		Course{CourseId: "2", CourseName: "node", CoursePrice: 499, Author: &Author{Fullname: "Ethel Greene", Website: "http://nose.lr/bavojjev"}},
		Course{CourseId: "3", CourseName: "golang", CoursePrice: 499, Author: &Author{Fullname: "Matilda Wolfe", Website: "http://noom.lu/pohri"}},
		Course{CourseId: "4", CourseName: "angular", CoursePrice: 499, Author: &Author{Fullname: "Lily Becker", Website: "http://zakaba.tg/opjot"}},
	)

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/create", createOneCourse).Methods("POST")
	r.HandleFunc("/update", updateOneCourse).Methods("PATCH")
	r.HandleFunc("/delete", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controller - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my first API in golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("Content-type", "application/json")

	// encode the json and send it to the req maker
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-type", "application/json")

	// grab id from req
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	return
}

func createOneCourse (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data found")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data found")
		return
	}

	for _, courseOld := range courses {
		fmt.Println(course)
		if courseOld.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("A course with same name already exists")
			return
		}
	}

	// generate a unique id, convert to string
	// append course to the db

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-type", "application/json")

	// grab id from req
	params := mux.Vars(r)

	// loop though, when id hit, remove id from array, add with id from req
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a res when id is not found
	
}

func deleteOneCourse (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// grab id from req
	params := mux.Vars(r)

	// loop though, when id hit, remove id from array, add with id from req
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course removed")
			return
		}
	}
}