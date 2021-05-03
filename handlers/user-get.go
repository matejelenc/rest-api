package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

// swagger:route GET /users users getUsers
// Returns a list of users
// responses:
//	200: usersResponse
//  400: badRequestResponse

//GetUsers returns all users from the database
func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	//validate the jwt token
	_, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	//get all users
	var people []data.Person
	foundPeople := data.DB.Find(&people)
	if foundPeople.Error != nil {
		http.Error(w, "No users exist", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&people)
}

// swagger:route GET /users/{id} users getUser
// Returns a user
// responses:
//	200: usersResponse
//  400: badRequestResponse

//GetPerson returns a user
func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	var person data.Person

	//validate the jwt token
	_, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	//get the user
	foundPerson := data.DB.First(&person, params["id"])
	if foundPerson.Error != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(person)
}
