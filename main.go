package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

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

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {

	courses = append(courses, Course{CourseId: "1", CourseName: "React", CoursePrice: 200, Author: &Author{Fullname: "telusko", Website: "telusko.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "angular", CoursePrice: 400, Author: &Author{Fullname: "lco", Website: "lco.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "node js", CoursePrice: 600, Author: &Author{Fullname: "barc", Website: "barc.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "python", CoursePrice: 700, Author: &Author{Fullname: "ak", Website: "ak.com"}})

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourses).Methods("GET")
	r.HandleFunc("/CreateCourses", CreateOneCourses).Methods("POST")
	r.HandleFunc("/UpdateCourses/{id}", UpdateCourses).Methods("PUT")
	r.HandleFunc("/deleteOneCourses/{id}", deleteOneCourses).Methods("DELETE")

	http.ListenAndServe(":7000", r)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to cred opertaion of api<h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getonecourse")
	w.Header().Set("Content-Type", "applicatioan/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("There no such course:)")
	return
}

func CreateOneCourses(w http.ResponseWriter, r *http.Request) {
	var course Course
	w.Header().Set("Content-Type", "applicatioan/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please enter the course details")
	}
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	rand.Seed(time.Now().UnixMicro())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}

func UpdateCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	params := mux.Vars(r)
	var course Course
	for index, courser := range courses {
		if courser.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			courser.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courser)
			return
		}
	}
}

func deleteOneCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, cours := range courses {
		if cours.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(courses)
			break
		}
	}
	return
}
