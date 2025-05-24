package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range students {
		if item.ID == params["id"] {
			students = append(students[:i], students[i+1:]...)
			var updated Student
			_ = json.NewDecoder(r.Body).Decode(&updated)
			updated.ID = params["id"]
			students = append(students, updated)
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range students {
		if item.ID == params["id"] {
			students = append(students[:i], students[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}

func main() {
	router := mux.NewRouter()
	students = append(students, Student{ID: "1", Name: "Anurag", Email: "anurag@example.com"})

	router.HandleFunc("/students", getStudents).Methods("GET")
	router.HandleFunc("/students/{id}", getStudent).Methods("GET")
	router.HandleFunc("/students", createStudent).Methods("POST")
	router.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
