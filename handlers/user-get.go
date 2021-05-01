package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
)

// swagger:route GET /users users getUsers
// Returns a list of users
// responses:
//	200: usersResponse

//GEtUsers returns all users from the database
func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var people []data.Person
	foundPeople := DB.Find(&people)
	if foundPeople.Error != nil {
		http.Error(w, "No users exist", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	var person data.Person

	foundPerson := DB.First(&person, params["id"])
	if foundPerson.Error != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(person)
}
