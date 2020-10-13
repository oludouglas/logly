package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	ID   string
	Name string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/users":
		users := []User{{ID: "1"}, {ID: "2"}}
		json.NewEncoder(w).Encode(&users)

	default:
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		got := User{ID: id}
		json.NewEncoder(w).Encode(&got)
	}
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodPut:
		UpdateUser(w, r)
	case http.MethodGet:
		GetUsers(w, r)
	}
}
