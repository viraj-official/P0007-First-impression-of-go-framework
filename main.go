package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve user data from userArray or any other data source
	jsonData, err := json.Marshal(userArray)
	if err != nil {
		http.Error(w, "Failed to retrieve user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsonData)
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var createUserReq CreateUserRequest
	err = json.Unmarshal(requestBody, &createUserReq)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	newUser := User{
		ID:   len(userArray) + 1,
		Name: createUserReq.Name,
		Age:  createUserReq.Age,
	}

	userArray = append(userArray, newUser)

	jsonData, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, "Failed to serialize user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write(jsonData)
}

var userArray []User

func main() {
	fmt.Println("Server up and running:8080")
	http.HandleFunc("/users", getUserHandler)
	http.HandleFunc("/users/create", createUserHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
