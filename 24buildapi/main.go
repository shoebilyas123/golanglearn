package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}


type Author struct {
	FullName string `json:"fullname"`
	Website string `json:"website"`
}


var courses []Course


// middleware - helper file
func isEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main(){
	r := mux.NewRouter();

	// Seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", Price: 240, Author: &Author{FullName: "MyNane",Website: "lco.dev"}},Course{CourseId: "1232", CourseName: "Golang", Price: 140, Author: &Author{FullName: "Shoeb",Website: "shoebilyas.com"}})

	r.HandleFunc("/", serveHome)
	r.HandleFunc("/courses/create",createCourse).Methods("POST")
	r.HandleFunc("/courses/get/all",getAllCourses).Methods("GET")
	r.HandleFunc("/courses/get/{id}",getCourse)

	log.Fatal(http.ListenAndServe(":4000",r))

}


// API: 
//  - Courses - CRUD with verifiying with schema
// 	- 

// Controllers
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcom To API. Please use testing routes to test</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)

}

func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","applicaction/json")
	// grab id from request
	params := mux.Vars(r)
	log.Printf("%v",params)

	// loop through courses and find the matching id and return
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")

}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)

	if err == io.EOF {
		json.NewEncoder(w).Encode("Please provide the course details")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId =  strconv.Itoa(rand.Intn(1000))
	
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	var updatePayload map[string]string
	err := json.NewDecoder(r.Body).Decode(&updatePayload)

	if err != nil {
		log.Printf("%v",err)
		json.NewEncoder(w).Encode("internal server error")
	}

	for index, course := range courses {
		if course.CourseId == updatePayload["id"] {
			courses = append(courses[:index],courses[index+1:]...)
			var c Course

			json.NewDecoder(r.Body).Decode(&c)
			c.CourseId = course.CourseId
			courses = append(courses, c)
			json.NewEncoder(w).Encode(c)
			return; 
		}
	}

	json.NewEncoder(w).Encode("course with given id not found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r);
	var foundAndDeleted bool = false;

	for index, course := range courses {
		if(course.CourseId == params["id"]) {
			courses = append(courses[:index], courses[index+1:]...)
			foundAndDeleted = true;
			break;
		}
	}

	if foundAndDeleted {
		json.NewEncoder(w).Encode("Course updated successfully")
	} else {
		
		json.NewEncoder(w).Encode("Course with the given id not found")
	}

	
}